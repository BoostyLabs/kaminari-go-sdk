package client

import (
	"fmt"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type createOnChainInvoiceResp struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

func (c *client) createOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*createOnChainInvoiceResp, error) {
	url := fmt.Sprintf("%v/invoice", c.onChainURL())
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

func (c *client) createLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	url := fmt.Sprintf("%v/invoice", c.lightningURL())
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
