package repository

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"

	"github.com/avran02/medods/config"
)

type Postgres interface {
	GetRefreshTokenInfo(userID string) (refreshTokenHash, accessTokenID string, err error)
	SaveNewRefreshToken(userID, refreshTokenHash, accessTokenID string) error
	GetUserEmail(userID string) (string, error)

	ClosePostgresConnection() error
}

type postgres struct {
	db *sql.DB
}

func (p *postgres) GetRefreshTokenInfo(userID string) (refreshTokenHash, accessTokenID string, err error) {
	slog.Info("Postgres.GetRefreshTokenInfo", "userID", userID)

	query := `SELECT token_hash, access_token_id 
	FROM refresh_tokens 
	WHERE user_id = $1
	`

	err = p.db.QueryRow(query, userID).Scan(&refreshTokenHash, &accessTokenID)
	if err != nil {
		return "", "", fmt.Errorf("can't get refresh token info: %w", err)
	}

	return refreshTokenHash, accessTokenID, nil
}

func (p *postgres) SaveNewRefreshToken(userID, refreshTokenHash, accessTokenID string) error {
	slog.Info("Postgres.SaveNewRefreshToken")

	slog.Debug("Postgres.SaveNewRefreshToken", "userID", userID, "refreshTokenHash", refreshTokenHash, "accessTokenID", accessTokenID)

	query := `INSERT INTO refresh_tokens (user_id, token_hash, access_token_id)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id)
	DO UPDATE SET 
		user_id = $1,
		token_hash = $2,
		access_token_id = $3
	`

	if _, err := p.db.Exec(query, userID, refreshTokenHash, accessTokenID); err != nil {
		return fmt.Errorf("can't save new refresh token: %w", err)
	}

	return nil
}

func (p *postgres) GetUserEmail(userID string) (string, error) {
	slog.Info("Postgres.GetUserEmail", "userID", userID)

	query := `SELECT email FROM users WHERE id = $1`

	var email string
	if err := p.db.QueryRow(query, userID).Scan(&email); err != nil {
		return "", fmt.Errorf("can't get user email: %w", err)
	}

	return email, nil
}

func (p *postgres) ClosePostgresConnection() error {
	return p.db.Close()
}

func NewPostgres(conf *config.DBConfig) Postgres {
	dsn := getDsn(*conf)
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("can't connect to db:\n", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("can't ping db:\n", err)
	}

	slog.Info("db connected")
	return &postgres{
		db: database,
	}
}

func getDsn(conf config.DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Password,
		conf.Name,
	)
}
