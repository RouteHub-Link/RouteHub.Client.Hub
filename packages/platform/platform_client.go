package platform

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"

	"github.com/redis/go-redis/v9"
)

var (
	keyPrefix = "platform:"
)

type PlatformClientService struct {
	redisClient *redis.Client
	// Deployed platform ID from the environment
	platformId string
	logger     *slog.Logger
}

func NewPlatformClientService(rc *redis.Client, logger *slog.Logger, platformId string) *PlatformClientService {
	return &PlatformClientService{
		redisClient: rc,
		logger:      logger,
		platformId:  platformId,
	}
}

func (pcs *PlatformClientService) GetPlatform(ctx context.Context) (*Platform, error) {
	if pcs.platformId == "" {
		return nil, errors.New(strings.Join([]string{"key is empty", "prefix is :", keyPrefix}, " "))
	}

	concatedKey := strings.Join([]string{keyPrefix, pcs.platformId}, "")

	platform, err := pcs.redisClient.Get(ctx, concatedKey).Result()
	if err == redis.Nil {
		pcs.logger.LogAttrs(ctx, slog.LevelInfo, "Platform not found in redis", slog.String("platformId", pcs.platformId))

		platform := MockPlatform()
		pcs.SetPlatform(ctx, &platform)

		pcs.logger.LogAttrs(ctx, slog.LevelInfo, "Platform setting as mock", slog.String("platformId", pcs.platformId))

		return &platform, nil
	}

	if err != nil {
		return nil, err
	}

	var p Platform
	err = json.Unmarshal([]byte(platform), &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (pcs *PlatformClientService) SetPlatform(ctx context.Context, p *Platform) error {
	pcs.logger.LogAttrs(ctx, slog.LevelDebug, "SetPlatform Called", slog.String("platformId", pcs.platformId))
	if pcs.platformId == "" {
		return errors.New(strings.Join([]string{"key is empty", "prefix is :", keyPrefix}, " "))
	}

	if p == nil {
		return errors.New("platform is nil")
	}

	pcs.logger.LogAttrs(ctx, slog.LevelDebug, "Setting platform", slog.String("platformId", pcs.platformId), slog.Any("platform", p))

	concatedKey := strings.Join([]string{keyPrefix, pcs.platformId}, "")

	pb, err := json.Marshal(p)
	if err != nil {
		return err
	}

	err = pcs.redisClient.Set(ctx, concatedKey, pb, 0).Err()
	if err != nil {
		return err
	}

	pcs.logger.Log(ctx, slog.LevelDebug, "Platform set", slog.String("platformId", pcs.platformId), slog.Any("platform", p))
	return nil
}
