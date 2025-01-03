package platform

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
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

func NewPlatformClientService(rc *redis.Client, logger *slog.Logger, platformId string, seed bool) *PlatformClientService {
	ctx := context.Background()

	pcs := &PlatformClientService{
		redisClient: rc,
		logger:      logger,
		platformId:  platformId,
	}

	if seed {
		pcs.clearPlatforms(ctx)
		mockPlatform := MockPlatform()
		pcs.SetPlatform(ctx, &mockPlatform)
	}

	return pcs
}

func NewPlatformClientServiceDefault(rc *redis.Client, logger *slog.Logger) *PlatformClientService {
	platformId := os.Getenv("PLATFORM_ID")
	if platformId == "" {
		platformId = "default"
		log.Println("PLATFORM_ID is not set, using default")
	}
	return NewPlatformClientService(rc, logger, platformId, false)
}

func (pcs *PlatformClientService) GetPlatform(ctx context.Context) (*Platform, error) {
	if pcs.platformId == "" {
		pcs.platformId = "default"
		pcs.logger.LogAttrs(ctx, slog.LevelDebug, "platformId is empty, using default", slog.String("platformId", pcs.platformId))
	}

	concatedKey := strings.Join([]string{keyPrefix, pcs.platformId}, "")

	platform, err := pcs.redisClient.Get(ctx, concatedKey).Result()

	if err != nil {
		log.Println(fmt.Sprintf("Error getting platform : key %s", concatedKey), err)
		return nil, err
	}

	var p Platform
	err = json.Unmarshal([]byte(platform), &p)
	if err != nil {
		log.Println("Error unmarshalling platform", err)
		return nil, err
	}

	return &p, nil
}

func (pcs *PlatformClientService) SetPlatform(ctx context.Context, p *Platform) error {
	pcs.logger.LogAttrs(ctx, slog.LevelDebug, "SetPlatform Called", slog.String("platformId", pcs.platformId))
	if pcs.platformId == "" {
		pcs.platformId = "default"
		pcs.logger.LogAttrs(ctx, slog.LevelDebug, "platformId is empty, using default", slog.String("platformId", pcs.platformId))
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

func (pcs *PlatformClientService) clearPlatforms(ctx context.Context) error {
	pcs.logger.Log(ctx, slog.LevelDebug, "Clearing Platforms")
	keys, err := pcs.redisClient.Keys(ctx, keyPrefix+"*").Result()
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		return nil
	}

	err = pcs.redisClient.Del(ctx, keys...).Err()
	if err != nil {
		return err
	}

	return nil
}
