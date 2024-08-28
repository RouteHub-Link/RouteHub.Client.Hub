package handlers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/RouteHub-Link/routehub.client.hub/packages/platform"
)

func (lh *PlatformHandlers) Set(payload []byte) error {
	ctx := context.Background()
	_Platform := &platform.Platform{}

	err := json.Unmarshal(payload, _Platform)
	if err != nil {
		return err
	}

	PlatformService := platform.NewPlatformClientServiceDefault(lh.redisClient, lh.logger)
	err = PlatformService.SetPlatform(ctx, _Platform)

	return err
}

func (lh *PlatformHandlers) Delete(payload []byte) error {
	return errors.New("not implemented")
}

func (lh *PlatformHandlers) Get(payload []byte) ([]byte, error) {
	ctx := context.Background()

	PlatformService := platform.NewPlatformClientServiceDefault(lh.redisClient, lh.logger)
	_Platform, err := PlatformService.GetPlatform(ctx)
	if err != nil {
		return nil, err
	}

	PlatformJson, err := json.Marshal(_Platform)
	if err != nil {
		return nil, err
	}

	return PlatformJson, nil
}

func (lh *PlatformHandlers) Fetch(payload []byte) ([]byte, error) {
	return nil, errors.New("not implemented")
}
