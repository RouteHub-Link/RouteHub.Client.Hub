package link

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"strings"

	"github.com/redis/go-redis/v9"
)

// This is a Link Client Service read from clients/redis.go
// Gets a key and checks if it exists in the redis database

type LinkClientService struct {
	redisClient *redis.Client
	logger      *slog.Logger
}

const (
	keyPrefix = "link:"
)

func NewLinkClientService(rc *redis.Client, logger *slog.Logger) *LinkClientService {
	return &LinkClientService{
		redisClient: rc,
		logger:      logger,
	}
}

func (lcs *LinkClientService) GetLink(ctx context.Context, key string) (*Link, error) {
	if key == "" {
		return nil, errors.New(strings.Join([]string{"key is empty", "prefix is :", keyPrefix}, " "))
	}

	concatedKey := strings.Join([]string{keyPrefix, key}, "")

	linkJson, err := lcs.redisClient.Get(ctx, concatedKey).Result()

	if err != nil {
		return nil, err
	}

	var link *Link
	err = json.Unmarshal([]byte(linkJson), &link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (lcs *LinkClientService) SetLink(ctx context.Context, link *Link) error {
	if link == nil {
		return errors.New("link is nil")
	}

	linkJson, err := json.Marshal(link)
	if err != nil {
		return err
	}

	concatedKey := strings.Join([]string{keyPrefix, link.ID}, "")

	err = lcs.redisClient.Set(ctx, concatedKey, linkJson, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
