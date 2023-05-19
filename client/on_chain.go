package client

import (
	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type createOnChainInvoiceResp struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

func (c *Client) createOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*createOnChainInvoiceResp, error) {
	url := "/api/bitcoin/v1/invoice"
	var result createOnChainInvoiceResp
	resp, err := c.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}

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
