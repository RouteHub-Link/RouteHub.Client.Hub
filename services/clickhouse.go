package services

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var (
	onceClickhouseClient sync.Once
	ClickhouseClient     *driver.Conn
)

func NewClickhouseClient(ctx context.Context, cc ClickhouseConfig, dc DetailsConfig) (clickhouseConnection *driver.Conn, err error) {
	onceClickhouseClient.Do(func() {
		conn, _err := connect(cc, dc)
		if _err != nil {
			logger.Log(ctx, slog.LevelWarn, "Error connecting to clickhouse", slog.String("error", _err.Error()))
			err = _err
		} else {
			logger.Log(ctx, slog.LevelInfo, "Clickhouse connected", slog.String("addr", strings.Join([]string{cc.Host, cc.Port}, ":")))
		}

		ClickhouseClient = &conn
	})

	if ClickhouseClient == nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return ClickhouseClient, nil
}

func connect(cc ClickhouseConfig, dc DetailsConfig) (driver.Conn, error) {
	var (
		ctx       = context.Background()
		conn, err = clickhouse.Open(&clickhouse.Options{
			Addr: []string{strings.Join([]string{cc.Host, cc.Port}, ":")},
			Auth: clickhouse.Auth{
				Database: cc.Database,
				Username: cc.Username,
				Password: cc.Password,
			},
			ClientInfo: clickhouse.ClientInfo{
				Products: []struct {
					Name    string
					Version string
				}{
					{Name: dc.Name, Version: dc.Version},
				},
			},
			DialTimeout:          time.Second * 30,
			MaxOpenConns:         5,
			MaxIdleConns:         5,
			ConnMaxLifetime:      time.Duration(10) * time.Minute,
			ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
			BlockBufferSize:      10,
			MaxCompressionBuffer: 10240,
			Debugf: func(format string, v ...interface{}) {
				logger.Log(ctx, slog.LevelDebug, "Clickhouse Debug", slog.String("format", fmt.Sprintf(format, v...)))
			},
			TLS: nil,
		})
	)

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			logger.Log(ctx, slog.LevelWarn, "Clickhouse Exception", slog.String("message", exception.Message), slog.String("stacktrace", exception.StackTrace))
		}
		return nil, err
	}

	return conn, nil
}

func GetClickhouseClient() *driver.Conn {
	if ClickhouseClient == nil {
		logger.Log(context.Background(), slog.LevelWarn, "Clickhouse client is nil")
	}

	return ClickhouseClient
}
