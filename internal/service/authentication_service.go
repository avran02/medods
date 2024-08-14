package service

import (
	"fmt"
	"log/slog"

	"github.com/avran02/medods/internal/pkg/jwt"
	"github.com/avran02/medods/internal/repository"
	"github.com/avran02/medods/internal/utils"
)

type AuthenticationService interface {
	RefreshTokens(refreshToken, userIP string) (newAccessToken, newRefreshToken string, err error)
	GetTokens(userID, userIP string) (accessToken, refreshToken string, err error)
	CheckIPChanged(token, userIP string) (userID string, IPChanged bool, err error)
}

type authenticationService struct {
	r   repository.Postgres
	jwt jwt.JwtGenerator
}

func (s *authenticationService) RefreshTokens(refreshTokenStr, userIP string) (newAccessToken, newRefreshToken string, err error) {
	slog.Info("authenticationService.RefreshTokens")
	refreshToken, err := s.jwt.ParseRefreshToken(refreshTokenStr)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.RefreshTokens: can't validate refresh token: %w", err)
	}

	writtenRefreshTokenHash, writtenAccessTokenID, err := s.r.GetRefreshTokenInfo(refreshToken.Subject)
	slog.Debug("authenticationService.RefreshTokens", "writtenRefreshTokenHash", writtenRefreshTokenHash, "writtenAccessTokenID", writtenAccessTokenID)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.RefreshTokens: can't get refresh token info: %w", err)
	}

	if err := utils.CompareHashAndPassword(refreshTokenStr, writtenRefreshTokenHash); err != nil {
		return "", "", fmt.Errorf("wrong refresh token: %w", err)
	}

	if refreshToken.AccessTokenID != writtenAccessTokenID {
		slog.Error("wrong access token id", "writtenAccessTokenID", writtenAccessTokenID, "refreshToken.AccessTokenID", refreshToken.AccessTokenID)
		return "", "", fmt.Errorf("wrong access token id: %w", ErrWrongTokensPair)
	}

	newAccessToken, newAccessTokenID, newRefreshToken, err := s.jwt.Generate(refreshToken.Subject, userIP)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.RefreshTokens: can't generate new tokens: %w", err)
	}

	newRefreshTokenHash, err := utils.Hash(newRefreshToken)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.RefreshTokens: can't hash new access token: %w", err)
	}

	if err := s.r.SaveNewRefreshToken(refreshToken.Subject, newRefreshTokenHash, newAccessTokenID); err != nil {
		return "", "", fmt.Errorf("authenticationService.RefreshTokens: can't save new tokens: %w", err)
	}

	return newAccessToken, newRefreshToken, nil
}

func (s *authenticationService) CheckIPChanged(token, userIP string) (string, bool, error) {
	slog.Info("authenticationService.CheckIPChanged")
	refreshToken, err := s.jwt.ParseRefreshToken(token)
	if err != nil {
		return "", false, fmt.Errorf("authenticationService.CheckIPChanged: can't validate refresh token: %w", err)
	}

	if refreshToken.UserIP != userIP {
		return refreshToken.Subject, true, nil
	}

	return refreshToken.Subject, false, nil
}

func (s *authenticationService) GetTokens(userID, userIP string) (accessToken, refreshToken string, err error) {
	slog.Info("authenticationService.GetTokens")

	accessToken, accessTokenID, refreshToken, err := s.jwt.Generate(userID, userIP)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.GetTokens: can't generate tokens: %w", err)
	}

	refreshTokenHash, err := utils.Hash(refreshToken)
	if err != nil {
		return "", "", fmt.Errorf("authenticationService.GetTokens: can't hash refresh token: %w", err)
	}

	if err := s.r.SaveNewRefreshToken(userID, refreshTokenHash, accessTokenID); err != nil {
		return "", "", fmt.Errorf("authenticationService.GetTokens: can't save refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

func NewAuthenticationService(r repository.Postgres, j jwt.JwtGenerator) AuthenticationService {
	return &authenticationService{
		r:   r,
		jwt: j,
	}
}
