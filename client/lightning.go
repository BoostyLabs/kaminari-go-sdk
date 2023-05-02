package client

import (
	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (c *Client) createLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	url := "/api/lightning/v1/invoice"
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
	url := "/api/lightning/v1/payment/send"
	resp, err := c.restyClient.R().
		SetBody(req).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return err
	}

	return nil
}
