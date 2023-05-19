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
		cfg:         cfg,
		restyClient: restyClient,
	}
}

func (c *Client) CreateOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (string, error) {
	resp, err := c.createOnChainInvoice(req)
	if err != nil {
		return "", errors.Wrap(err, "can't create on-chain invoice")
	}

	return resp.BitcoinAddress, nil
}

func (c *Client) CreateLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	resp, err := c.createLightningInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't create lightning invoice")
	}

	return resp, nil
}

func (c *Client) SendOnChainPayment(req *kaminarigosdk.SendOnChainPaymentRequest) error {
	err := c.sendOnChainPayment(req)
	if err != nil {
		return errors.Wrap(err, "can't send on-chain payment")
	}

	return nil
}

func (c *Client) SendLightningPayment(req *kaminarigosdk.SendLightningPaymentRequest) error {
	err := c.sendLightningPayment(req)
	if err != nil {
		return errors.Wrap(err, "can't send lightning payment")
	}

	return nil
}

func (c *Client) GetOnChainInvoice(req *kaminarigosdk.GetOnChainInvoiceRequest) (*kaminarigosdk.GetOnChainInvoiceResponse, error) {
	resp, err := c.getOnChainInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain invoice")
	}

	return resp, nil
}

func (c *Client) GetLightningInvoice(req *kaminarigosdk.GetLightningInvoiceRequest) (*kaminarigosdk.GetLightningInvoiceResponse, error) {
	resp, err := c.getLightningInvoice(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning invoice")
	}

	return resp, nil
}

func (c *Client) GetOnChainTransaction(req *kaminarigosdk.GetOnChainTransactionRequest) (*kaminarigosdk.GetOnChainTransactionResponse, error) {
	resp, err := c.getOnChainTransaction(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain transaction")
	}

	return resp, nil
}

func (c *Client) GetLightningTransaction(req *kaminarigosdk.GetLightningTransactionRequest) (*kaminarigosdk.GetLightningTransactionResponse, error) {
	resp, err := c.getLightningTransaction(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return resp, nil
}

func (c *Client) VerifyWebhookSignature(req *kaminarigosdk.VerifyWebhookSignatureRequest) (*kaminarigosdk.VerifyWebhookSignatureResponse, error) {
	resp, err := c.verifyWebhookSignature(req)
	if err != nil {
		return nil, errors.Wrap(err, "can't verify webhook signature")
	}

	return resp, nil
}
