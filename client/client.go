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

func DefaultClient(cfg *Config) kaminarigosdk.Interface {
	if !isCfgValid(cfg) {
		err := errors.Errorf("kaminari config is not valid")
		log.Error(err)
		return nil
	}

	restyClient := resty.New().
		SetRetryCount(3).
		SetHeader("X-kaminari-api-key", cfg.GetApiKey()).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(15 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r != nil && r.StatusCode() == http.StatusTooManyRequests
			},
		)

	return &Client{
		restyClient: restyClient,
	}
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

func (a *Client) SendOnChainPayment(req *kaminarigosdk.SendOnChainPaymentRequest) error {
	err := a.sendOnChainPayment(req)
	if err != nil {
		return errors.Wrap(err, "can't send on-chain payment")
	}

	return nil
}

func (a *Client) SendLightningPayment(req *kaminarigosdk.SendLightningPaymentRequest) error {
	err := a.sendLightningPayment(req)
	if err != nil {
		return errors.Wrap(err, "can't send lightning payment")
	}

	return nil
}

func (a *Client) GetOnChainInvoice(req *kaminarigosdk.GetOnChainInvoiceRequest) (*kaminarigosdk.GetOnChainInvoiceResponse, error) {
	resp, err := a.getOnChainInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain invoice")
	}

	return resp, nil
}

func (a *Client) GetLightningInvoice(req *kaminarigosdk.GetLightningInvoiceRequest) (*kaminarigosdk.GetLightningInvoiceResponse, error) {
	resp, err := a.getLightningInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning invoice")
	}

	return resp, nil
}

func (a *Client) GetOnChainTransaction(req *kaminarigosdk.GetOnChainTransactionRequest) (*kaminarigosdk.GetOnChainTransactionResponse, error) {
	resp, err := a.getOnChainTransaction(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain transaction")
	}

	return resp, nil
}

func (a *Client) GetLightningTransaction(req *kaminarigosdk.GetLightningTransactionRequest) (*kaminarigosdk.GetLightningTransactionResponse, error) {
	resp, err := a.getLightningTransaction(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return resp, nil
}
