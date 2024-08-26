package link

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/redis/go-redis/v9"
)

// This is a Link Client Service read from clients/redis.go
// Gets a key and checks if it exists in the redis database

type LinkClientService struct {
	redisClient *redis.Client
}

const (
	LinkKeyPrefix = "link:"
)

func NewLinkClientService(rc *redis.Client) *LinkClientService {
	return &LinkClientService{
		redisClient: rc,
	}
}

func (lcs *LinkClientService) GetLink(key string) (link *Link, err error) {
	if key == "" {
		return nil, errors.New("key is empty")
	}

	ctx := context.Background()

	concatedKey := strings.Join([]string{LinkKeyPrefix, key}, "")

	linkJson, err := lcs.redisClient.Get(ctx, concatedKey).Result()

	if err != nil {
		return nil, err
	}

	link = &Link{}
	err = json.Unmarshal([]byte(linkJson), &link)
	if err != nil {
		return nil, err
	}

	return link, nil
}

func (lcs *LinkClientService) SetLink(link *Link) error {
	if link == nil {
		return errors.New("link is nil")
	}

	ctx := context.Background()

	linkJson, err := json.Marshal(link)
	if err != nil {
		return err
	}

	concatedKey := strings.Join([]string{LinkKeyPrefix, link.ID}, "")

	err = lcs.redisClient.Set(ctx, concatedKey, linkJson, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
