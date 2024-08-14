package config

import (
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	JWTConfig
	DBConfig
	ServerConfig
	SMTPConfig
}

type SMTPConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type JWTConfig struct {
	Secret         string
	AccessExpTime  int
	RefreshExpTime int
}

type ServerConfig struct {
	Host      string
	Port      string
	APIPrefix string
	LogLevel  string
}

func New() *Config {
	// if err := gotenv.Load(); err != nil {
	// 	log.Fatal("can't load .env file")
	// }

	AccessExpTimeStr := os.Getenv("ACCESS_EXP_TIME")
	RefreshExpTimeStr := os.Getenv("REFRESH_EXP_TIME")

	AccessExpTime, err := strconv.Atoi(AccessExpTimeStr)
	if err != nil {
		slog.Warn("can't parse ACCESS_EXP_TIME. Using default value", "value", DEFAULUT_ACCESS_EXP_TIME)
		AccessExpTime = DEFAULUT_ACCESS_EXP_TIME
	}

	RefreshExpTime, err := strconv.Atoi(RefreshExpTimeStr)
	if err != nil {
		slog.Warn("can't parse REFRESH_EXP_TIME. Using default value", "value", DEFAULUT_REFRESH_EXP_TIME)
		RefreshExpTime = DEFAULUT_REFRESH_EXP_TIME
	}

	SMTPPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		slog.Error("can't parse SMTP_PORT. Using default value", "value", DEFAULUT_SMTP_PORT)
	}

	return &Config{
		JWTConfig: JWTConfig{
			Secret:         os.Getenv("JWT_SECRET"),
			AccessExpTime:  AccessExpTime,
			RefreshExpTime: RefreshExpTime,
		},
		DBConfig: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		ServerConfig: ServerConfig{
			Host:      os.Getenv("SERVER_HOST"),
			Port:      os.Getenv("SERVER_PORT"),
			APIPrefix: os.Getenv("API_PREFIX"),
			LogLevel:  os.Getenv("LOG_LEVEL"),
		},
		SMTPConfig: SMTPConfig{
			Host:     os.Getenv("SMTP_HOST"),
			Port:     SMTPPort,
			User:     os.Getenv("SMTP_USER"),
			Password: os.Getenv("SMTP_PASSWORD"),
		},
	}
}
