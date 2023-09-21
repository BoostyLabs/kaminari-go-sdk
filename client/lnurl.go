package client

import (
	"fmt"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (c *Client) GetLightningAddress() (*kaminarigosdk.GetLightningAddrResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/lnurl/address", c.cfg.ApiUrl)
	var result kaminarigosdk.GetLightningAddrResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return &result, nil
}

func (c *Client) GetLightningAddressForMerchant(req *kaminarigosdk.GetLightningAddrForMerchantRequest) (*kaminarigosdk.GetLightningAddrForMerchantResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/lnurl/address/for/merchant/%v", c.cfg.ApiUrl, req.MerchantID)
	var result kaminarigosdk.GetLightningAddrForMerchantResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return &result, nil
}

func (c *Client) ConvertLnUrlInvoiceToLND(req *kaminarigosdk.ConvertLnUrlInvoiceToLNDRequest) (*kaminarigosdk.ConvertLnUrlInvoiceToLNDResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/invoice/from/lnurl/%v?amount=%v", c.cfg.ApiUrl, req.LnrulInvoice, req.Amount)
	var result kaminarigosdk.ConvertLnUrlInvoiceToLNDResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't convert ln invoice from lnurl")
	}

	return &result, nil
}
