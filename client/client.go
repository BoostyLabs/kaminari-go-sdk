package client

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type Client struct {
	cfg         *Config
	restyClient *resty.Client
}

func DefaultClient(cfg *Config) (kaminarigosdk.Interface, error) {
	if !isCfgValid(cfg) {
		err := errors.Errorf("kaminari config is not valid")
		log.Error(err)
		return nil, err
	}

	restyClient := resty.New().
		SetRetryCount(3).
		SetHeader("X-kaminari-api-key", cfg.ApiKey).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(15 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r != nil && r.StatusCode() == http.StatusTooManyRequests
			},
		)

	return &Client{
		cfg:         cfg,
		restyClient: restyClient,
	}, nil
}
