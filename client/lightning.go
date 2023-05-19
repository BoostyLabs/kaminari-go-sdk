package client

import (
	"fmt"
	"strconv"
	"time"

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

func (c *Client) getLightningInvoice(req *kaminarigosdk.GetLightningInvoiceRequest) (*kaminarigosdk.GetLightningInvoiceResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/invoices/%s", c.cfg.ApiUrl, req.ID)
	var result getLightningInvoiceResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	pbInvoice, err := toPbLightningInvoice(result.Invoice)
	if err != nil {
		return nil, err
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

func (c *Client) getLightningTransaction(req *kaminarigosdk.GetLightningTransactionRequest) (*kaminarigosdk.GetLightningTransactionResponse, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/transactions/%s", c.cfg.ApiUrl, req.ID)
	var result getLightningTransactionResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	pbTx, err := toPbLightningTx(result.Transaction)
	if err != nil {
		return nil, err
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
