package client

import (
	"fmt"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (c *Client) createLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/invoice", c.cfg.ApiUrl)
	var result kaminarigosdk.CreateLightningInvoiceResponse
	resp, err := c.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) sendLightningPayment(req *kaminarigosdk.SendLightningPaymentRequest) error {
	url := fmt.Sprintf("%s/api/lightning/v1/payment/send", c.cfg.ApiUrl)
	resp, err := c.restyClient.R().
		SetBody(req).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return err
	}

	return nil
}

func (c *Client) getLightningInvoice(req *kaminarigosdk.GetLightningInvoiceRequest) (*kaminarigosdk.GetLightningInvoiceResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/invoices/%s", c.cfg.ApiUrl, req.ID)
	var result kaminarigosdk.GetLightningInvoiceResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) getLightningTransaction(req *kaminarigosdk.GetLightningTransactionRequest) (*kaminarigosdk.GetLightningTransactionResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/transactions/%s", c.cfg.ApiUrl, req.ID)
	var result kaminarigosdk.GetLightningTransactionResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
