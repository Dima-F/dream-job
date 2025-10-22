package database

import (
	"context"

	"github.com/Dima-F/dream-job/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)

	if err != nil {
		logger.Error().Msg("Cant connect to db")
		panic(err)
	}
	logger.Info().Msg("Db connection successfull")

	return dbpool
}
