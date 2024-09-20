package services

import (
	"context"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	onceTimeScaleClient sync.Once
	TimeScaleClient     *pgxpool.Pool
)

func NewTimeScaleClient(ctx context.Context, cstr string) *pgxpool.Pool {
	onceTimeScaleClient.Do(func() {
		ctx := context.Background()
		connStr := cstr
		dbpool, err := pgxpool.New(ctx, connStr)
		if err != nil {
			logger.Log(ctx, slog.LevelError, "Error connecting to timescale", slog.String("error", err.Error()))
		}

		var greeting string
		err = dbpool.QueryRow(ctx, "select 'Hello, Timescale!'").Scan(&greeting)
		if err != nil {
			logger.Log(ctx, slog.LevelError, "Error reading from timescale", slog.String("error", err.Error()))
		}
		logger.Log(ctx, slog.LevelInfo, "Timescale greeting", slog.String("greeting", greeting))

		// TODO - Create tables if they don't exist

		TimeScaleClient = dbpool
	})

	return TimeScaleClient
}

func GetTimeScaleClient() *pgxpool.Pool {
	if TimeScaleClient == nil {
		logger.Log(context.Background(), slog.LevelError, "Redis client is nil")
	}

	return TimeScaleClient
}
