package analytics

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

const (
	MaxHeaderValueLength = 4096
)

type AnalyticsMiddleware struct {
	db                   *pgxpool.Pool
	logger               *slog.Logger
	maxHeaderValueLength int
}

type Option func(*AnalyticsMiddleware)

func WithLogger(logger *slog.Logger) Option {
	return func(am *AnalyticsMiddleware) {
		am.logger = logger
	}
}

func WithMaxHeaderValueLength(length int) Option {
	return func(am *AnalyticsMiddleware) {
		am.maxHeaderValueLength = length
	}
}

func NewAnalyticsMiddleware(db *pgxpool.Pool, opts ...Option) (*AnalyticsMiddleware, error) {
	am := &AnalyticsMiddleware{
		db:                   db,
		logger:               slog.Default(),
		maxHeaderValueLength: MaxHeaderValueLength,
	}

	for _, opt := range opts {
		opt(am)
	}

	if err := am.Setup(); err != nil {
		am.logger.Error("Failed to setup analytics tables", "error", err)
		return nil, err
	}

	return am, nil
}

func (am *AnalyticsMiddleware) Setup() error {
	_, err := am.db.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS request_logs (
			id SERIAL PRIMARY KEY,
			timestamp TIMESTAMPTZ NOT NULL,
			method TEXT NOT NULL,
			path TEXT NOT NULL,
			status INTEGER NOT NULL,
			ip_address TEXT NOT NULL,
			user_agent TEXT,
			referer TEXT,
			headers JSONB,
			response_time FLOAT
		);

		CREATE INDEX IF NOT EXISTS idx_request_logs_timestamp ON request_logs (timestamp);
	`)
	return err
}

func (am *AnalyticsMiddleware) Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)

			reqLog := RequestLog{
				Timestamp:    time.Now(),
				Method:       c.Request().Method,
				Path:         c.Request().URL.Path,
				Status:       c.Response().Status,
				IPAddress:    c.RealIP(),
				UserAgent:    truncateString(c.Request().UserAgent(), MaxHeaderValueLength),
				Referer:      truncateString(c.Request().Referer(), MaxHeaderValueLength),
				Headers:      getLimitedHeadersJSON(c.Request().Header),
				ResponseTime: time.Since(start).Seconds(),
			}

			if err := am.logRequest(c.Request().Context(), reqLog); err != nil {
				am.logger.Error("Failed to log request", "error", err)
			}

			return err
		}
	}
}

func (am *AnalyticsMiddleware) logRequest(ctx context.Context, log RequestLog) error {
	_, err := am.db.Exec(ctx, `
		INSERT INTO request_logs (timestamp, method, path, status, ip_address, user_agent, referer, headers, response_time)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`, log.Timestamp, log.Method, log.Path, log.Status, log.IPAddress, log.UserAgent, log.Referer, log.Headers, log.ResponseTime)

	if err != nil {
		am.logger.Error("Failed to insert log", "error", err)
	}

	am.logger.Debug("Logged request", "log", log)

	return err
}

type RequestLog struct {
	Timestamp    time.Time
	Method       string
	Path         string
	Status       int
	IPAddress    string
	UserAgent    string
	Referer      string
	Headers      json.RawMessage
	ResponseTime float64
}

func getLimitedHeadersJSON(headers http.Header) json.RawMessage {
	headerMap := make(map[string]string)
	for key, values := range headers {
		if len(values) > 0 {
			headerMap[key] = truncateString(values[0], MaxHeaderValueLength)
		}
	}
	jsonBytes, _ := json.Marshal(headerMap)
	return jsonBytes
}

func truncateString(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength]
}
