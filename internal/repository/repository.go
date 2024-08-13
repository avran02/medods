package repository

import "github.com/avran02/medods/config"

type Repository interface {
	Postgres
}

type repository struct {
	Postgres
}

func New(config *config.DBConfig) Repository {
	return &repository{
		Postgres: NewPostgres(config),
	}
}
