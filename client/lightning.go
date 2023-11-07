package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (client *Client) CreateLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	uriPath := "/api/lightning/v1/invoice"
	url := client.cfg.ApiUrl + uriPath

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	signature, err := client.GetSignature(uriPath, req.Nonce, body)
	if err != nil {
		return nil, err
	}

	var result kaminarigosdk.CreateLightningInvoiceResponse
	resp, err := client.restyClient.R().
		SetBody(req).
		SetResult(&result).
		SetHeader(ApiSignatureHeader, signature).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't create lightning invoice")
	}

	return &result, nil
}

func (client *Client) SendLightningPayment(req *kaminarigosdk.SendLightningPaymentRequest) error {
	uriPath := "/api/lightning/v1/payment/send"
	url := client.cfg.ApiUrl + uriPath

	body, err := json.Marshal(req)
	if err != nil {
		return err
	}

	signature, err := client.GetSignature(uriPath, req.Nonce, body)
	if err != nil {
		return err
	}

	resp, err := client.restyClient.R().
		SetBody(req).
		SetHeader(ApiSignatureHeader, signature).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return errors.Wrap(err, "can't send lightning payment")
	}

	return nil
}

type getLightningInvoiceResponse struct {
	Invoice *filteredLightningInvoice `json:"invoice"`
}

type filteredLightningInvoice struct {
	ID             string `json:"id"`
	EncodedInvoice string `json:"encodedInvoice"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}

func (client *Client) GetLightningInvoice(req *kaminarigosdk.GetLightningInvoiceRequest) (*kaminarigosdk.GetLightningInvoiceResponse, error) {
	uriPath := fmt.Sprintf("/api/lightning/v1/invoices/%s?nonce=%s", req.ID, req.Nonce)
	url := client.cfg.ApiUrl + uriPath

	var result getLightningInvoiceResponse

	signature, err := client.GetSignature(uriPath, req.Nonce, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.restyClient.R().
		SetResult(&result).
		SetHeader(ApiSignatureHeader, signature).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning invoice")
	}

	pbInvoice, err := toPbLightningInvoice(result.Invoice)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning invoice")
	}

	return &kaminarigosdk.GetLightningInvoiceResponse{
		Invoice: pbInvoice,
	}, nil
}

type getLightningTransactionResponse struct {
	Transaction *filteredLightningTransaction `json:"transaction"`
}

type filteredLightningTransaction struct {
	ID            string `json:"id"`
	MerchantID    string `json:"merchantId"`
	Status        string `json:"status"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	Amount        string `json:"amount"`
	CreatedAt     string `json:"createdAt"`
	Direction     string `json:"direction"`
	Confirmations string `json:"confirmations"`
	BlockNumber   string `json:"blockNumber"`
	ExplorerUrl   string `json:"explorerUrl"`
}

func (client *Client) GetLightningTransaction(req *kaminarigosdk.GetLightningTransactionRequest) (*kaminarigosdk.GetLightningTransactionResponse, error) {
	uriPath := fmt.Sprintf("/api/lightning/v1/transactions/%s?nonce=%s", req.ID, req.Nonce)
	url := client.cfg.ApiUrl + uriPath

	var result getLightningTransactionResponse

	signature, err := client.GetSignature(uriPath, req.Nonce, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.restyClient.R().
		SetHeader(ApiSignatureHeader, signature).
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	pbTx, err := toPbLightningTx(result.Transaction)
	if err != nil {
		return nil, errors.Wrap(err, "can't get lightning transaction")
	}

	return &kaminarigosdk.GetLightningTransactionResponse{
		Transaction: pbTx,
	}, nil
}

func toPbLightningInvoice(invoice *filteredLightningInvoice) (*kaminarigosdk.FilteredLightningInvoice, error) {
	amount, err := strconv.Atoi(invoice.Amount)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, invoice.CreatedAt)
	if err != nil {
		return nil, err
	}

	status, err := toPbInvoiceStatus(invoice.Status)
	if err != nil {
		return nil, err
	}

	return &kaminarigosdk.FilteredLightningInvoice{
		ID:             invoice.ID,
		EncodedInvoice: invoice.EncodedInvoice,
		Description:    invoice.Description,
		Amount:         int64(amount),
		Status:         status,
		CreatedAt: &kaminarigosdk.Timestamp{
			Seconds: int64(createdAt.Second()),
		},
	}, nil
}

func toPbLightningTx(tx *filteredLightningTransaction) (*kaminarigosdk.FilteredLightningTransaction, error) {
	amount, err := strconv.Atoi(tx.Amount)
	if err != nil {
		return nil, err
	}

	status, err := toPbTransactionStatus(tx.Status)
	if err != nil {
		return nil, err
	}

	createdAt, err := time.Parse(time.RFC3339, tx.CreatedAt)
	if err != nil {
		return nil, err
	}

	txType, err := toPbTransactionType(tx.Direction)
	if err != nil {
		return nil, err
	}

	return &kaminarigosdk.FilteredLightningTransaction{
		ID:          tx.ID,
		MerchantID:  tx.MerchantID,
		Status:      status,
		Source:      tx.Source,
		Destination: tx.Destination,
		Amount:      int64(amount),
		CreatedAt: &kaminarigosdk.Timestamp{
			Seconds: int64(createdAt.Second()),
		},
		Direction:   txType,
		ExplorerUrl: tx.ExplorerUrl,
	}, nil
}
