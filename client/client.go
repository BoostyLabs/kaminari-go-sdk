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
		restyClient: restyClient,
	}, nil
}

func (a *Client) CreateOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (string, error) {
	resp, err := a.createOnChainInvoice(req)
	if err != nil {
		return "", errors.Wrap(err, "can't create on-chain invoice")
	}

	return resp.BitcoinAddress, nil
}

func (a *Client) CreateLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	resp, err := a.createLightningInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't create lightning invoice")
	}

	return resp, nil
}
