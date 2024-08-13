package controller

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/avran02/medods/internal/dto"
	"github.com/avran02/medods/internal/service"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticationController interface {
	GetTokens(w http.ResponseWriter, r *http.Request)
	RefreshTokens(w http.ResponseWriter, r *http.Request)
}

type authenticationController struct {
	s service.AuthenticationService
}

func (c *authenticationController) GetTokens(w http.ResponseWriter, r *http.Request) {
	slog.Info("authenticationController.GetTokens")
	userID := r.URL.Query().Get("userID")
	userIP := getUserIP(r)

	accessToken, refreshToken, err := c.s.GetTokens(userID, userIP)
	if err != nil {
		slog.Error("authenticationController.GetTokens: can't get tokens", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.GetTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("authenticationController.GetTokens: can't encode response", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *authenticationController) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	slog.Info("authenticationController.RefreshTokens")
	var req dto.RefreshTokensRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("authenticationController.RefreshTokens: can't decode request", "err", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := c.s.RefreshTokens(req.RefreshToken, getUserIP(r))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			slog.Error("authenticationController.RefreshTokens: wrong refresh token", "err", err.Error())
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		slog.Error("authenticationController.RefreshTokens: can't refresh tokens", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := dto.RefreshTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		slog.Error("authenticationController.RefreshTokens: can't encode response", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUserIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	var addr string
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		slog.Info("X-Forwarded-For", "ips", ips)
		addr = ips[0]
	} else {
		addr = r.RemoteAddr
	}

	addrParts := strings.Split(addr, ":")
	ip := strings.Join(addrParts[:len(addrParts)-1], ":")
	slog.Info("RemoteAddr", "ip", ip)
	return ip
}

func NewAuthenticationController(s service.AuthenticationService) AuthenticationController {
	return &authenticationController{
		s: s,
	}
}
