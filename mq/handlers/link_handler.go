package handlers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/RouteHub-Link/routehub.client.hub/packages/link"
)

func (lh *LinkHandlers) Set(payload []byte) error {
	ctx := context.Background()

	_link := &link.Link{}

	err := json.Unmarshal(payload, _link)
	if err != nil {
		return err
	}

	linkService := link.NewLinkClientService(lh.redisClient, lh.logger)
	err = linkService.SetLink(ctx, _link)

	return err
}

func (lh *LinkHandlers) Delete(payload []byte) error {
	ctx := context.Background()
	key := string(payload)
	if key == "" {
		return errors.New("key is empty")
	}

	linkService := link.NewLinkClientService(lh.redisClient, lh.logger)
	err := linkService.DelLink(ctx, key)

	return err
}

func (lh *LinkHandlers) Get(payload []byte) ([]byte, error) {
	ctx := context.Background()
	key := string(payload)

	linkService := link.NewLinkClientService(lh.redisClient, lh.logger)
	_link, err := linkService.GetLink(ctx, key)
	if err != nil {
		return nil, err
	}

	linkJson, err := json.Marshal(_link)
	if err != nil {
		return nil, err
	}

	return linkJson, nil
}

func (lh *LinkHandlers) Fetch(payload []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
