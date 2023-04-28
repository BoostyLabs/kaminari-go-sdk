//go:build !mock && !scale
// +build !mock,!scale

package client

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type Adapter struct {
	c *client
}

func DefaultAdapter(cfg *Config) kaminarigosdk.Interface {
	if !isCfgValid(cfg) {
		err := errors.Errorf("kaminari config is not valid")
		log.Error(err)
		return NewNotConf()
	}

	apiKey, err := getSecretValue(cfg.GetApiKey())
	if err != nil {
		err := errors.Wrap(err, "can't get secret value from kaminari config")
		log.Error(err)
		return NewNotConf()
	}

	restyClient := resty.New().
		SetRetryCount(3).
		SetHeader("X-kaminari-api-key", apiKey).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(15 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r != nil && r.StatusCode() == http.StatusTooManyRequests
			},
		)
	client := newKaminariClient(cfg.GetApiUrl(), restyClient)

	return &Adapter{
		c: client,
	}
}

func (a *Adapter) CreateOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (string, error) {
	resp, err := a.c.createOnChainInvoice(req)
	if err != nil {
		return "", errors.Wrap(err, "can't create on-chain invoice")
	}

	return resp.BitcoinAddress, nil
}

func (a *Adapter) CreateLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	resp, err := a.c.createLightningInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't create lightning invoice")
	}

	return resp, nil
}
