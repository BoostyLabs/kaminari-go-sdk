package client

import (
	"fmt"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (client *Client) GetLightningAddress(nonce string) (*kaminarigosdk.GetLightningAddrResponse, error) {
	uriPath := fmt.Sprintf("/api/lightning/v1/lnurl/address?nonce=%s", nonce)
	url := fmt.Sprintf("%s/%s", client.cfg.ApiUrl, uriPath)

	var result kaminarigosdk.GetLightningAddrResponse

	signature, err := client.GetSignature(uriPath, nonce, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.restyClient.R().
		SetResult(&result).
		SetHeader(ApiSignatureHeader, signature).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return &result, nil
}

func (client *Client) GetLightningAddressForMerchant(req *kaminarigosdk.GetLightningAddrForMerchantRequest) (*kaminarigosdk.GetLightningAddrForMerchantResponse, error) {
	uriPath := fmt.Sprintf("/api/lightning/v1/lnurl/address/for/merchant/%s?nonce=%s", req.MerchantID, req.Nonce)
	url := fmt.Sprintf("%s/%s", client.cfg.ApiUrl, uriPath)

	var result kaminarigosdk.GetLightningAddrForMerchantResponse

	signature, err := client.GetSignature(uriPath, req.Nonce, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.restyClient.R().
		SetResult(&result).
		SetHeader(ApiSignatureHeader, signature).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return &result, nil
}

func (client *Client) ConvertLnUrlInvoiceToLND(req *kaminarigosdk.ConvertLnUrlInvoiceToLNDRequest) (*kaminarigosdk.ConvertLnUrlInvoiceToLNDResponse, error) {
	uriPath := fmt.Sprintf("/api/lightning/v1/invoice/from/lnurl/%v?amount=%v", req.LnrulInvoice, req.Amount)
	url := client.cfg.ApiUrl + uriPath

	var result kaminarigosdk.ConvertLnUrlInvoiceToLNDResponse

	signature, err := client.GetSignature(uriPath, req.Nonce, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.restyClient.R().
		SetResult(&result).
		SetHeader(ApiSignatureHeader, signature).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't convert ln invoice from lnurl")
	}

	return &result, nil
}
