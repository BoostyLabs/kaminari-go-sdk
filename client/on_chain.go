package client

import (
	"fmt"

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

func (c *Client) sendOnChainPayment(req *kaminarigosdk.SendOnChainPaymentRequest) error {
	url := "/api/bitcoin/v1/payment/send"
	resp, err := c.restyClient.R().
		SetBody(req).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return err
	}

	return nil
}

func (c *Client) getOnChainInvoice(req *kaminarigosdk.GetOnChainInvoiceRequest) (*kaminarigosdk.GetOnChainInvoiceResponse, error) {
	url := fmt.Sprintf("/api/bitcoin/v1/invoices/{%s}", req.BitcoinAddress)
	var result kaminarigosdk.GetOnChainInvoiceResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
