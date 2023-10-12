package client

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type createOnChainInvoiceResp struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

func (c *Client) CreateOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (string, error) {
	url := fmt.Sprintf("%s/api/bitcoin/v1/invoice", c.cfg.ApiUrl)
	var result createOnChainInvoiceResp

	resp, err := c.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return "", errors.Wrap(err, "can't create on-chain invoice")
	}

	return result.BitcoinAddress, nil
}

func (c *Client) SendOnChainPayment(req *kaminarigosdk.SendOnChainPaymentRequest) error {
	url := fmt.Sprintf("%s/api/bitcoin/v1/payment/send", c.cfg.ApiUrl)
	resp, err := c.restyClient.R().
		SetBody(req).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return errors.Wrap(err, "can't send on-chain payment")
	}

	return nil
}

type getOnChainInvoiceResponse struct {
	Invoice *filteredOnChainInvoice `json:"invoice"`
}

type filteredOnChainInvoice struct {
	BitcoinAddress string `json:"bitcoinAddress"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}

func (c *Client) GetOnChainInvoice(req *kaminarigosdk.GetOnChainInvoiceRequest) (*kaminarigosdk.GetOnChainInvoiceResponse, error) {
	url := fmt.Sprintf("%s/api/bitcoin/v1/invoices/%s", c.cfg.ApiUrl, req.BitcoinAddress)
	var result getOnChainInvoiceResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get on-chain invoice")
	}

	pbInvoice, err := toPbInvoice(result.Invoice)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain invoice")
	}

	return &kaminarigosdk.GetOnChainInvoiceResponse{
		Invoice: pbInvoice,
	}, nil
}

type getOnChainTransactionResponse struct {
	Transaction *filteredOnChainTransaction `json:"transaction"`
}

type filteredOnChainTransaction struct {
	ID            string `json:"id"`
	MerchantID    string `json:"merchantId"`
	Status        string `json:"status"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	Amount        string `json:"amount"`
	CreatedAt     string `json:"createdAt"`
	Direction     string `json:"direction"`
	Confirmations int32  `json:"confirmations"`
	BlockNumber   string `json:"blockNumber"`
	ExplorerUrl   string `json:"explorerUrl"`
}

func (c *Client) GetOnChainTransaction(req *kaminarigosdk.GetOnChainTransactionRequest) (*kaminarigosdk.GetOnChainTransactionResponse, error) {
	url := fmt.Sprintf("%s/api/bitcoin/v1/transactions/%s", c.cfg.ApiUrl, req.ID)
	var result getOnChainTransactionResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get on-chain transaction")
	}

	pbTx, err := toPbTx(result.Transaction)
	if err != nil {
		return nil, errors.Wrap(err, "can't get on-chain transaction")
	}

	return &kaminarigosdk.GetOnChainTransactionResponse{
		Transaction: pbTx,
	}, nil
}

type estimateOnChainTxResponse struct {
	Fee string `json:"fee"`
}

// EstimateOnChainTx estimates fee for on-chain tx, estimated fee returns in satoshi.
// Provided amount should be in satoshi(1 BTC = 100_000_000 sats).
func (c *Client) EstimateOnChainTx(req *kaminarigosdk.EstimateOnChainTxRequest) (*kaminarigosdk.EstimateOnChainTxResponse, error) {
	url := fmt.Sprintf("%s/api/bitcoin/v1/tx/estimate?bitcoin_address=%v&amount=%v", c.cfg.ApiUrl, req.BitcoinAddress, req.Amount)
	var result estimateOnChainTxResponse

	resp, err := c.restyClient.R().
		SetResult(&result).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get on-chain transaction")
	}

	estimateRes, err := toPbEstimateFee(&result)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse estimate fee response")
	}

	return estimateRes, nil
}

func toPbEstimateFee(estimate *estimateOnChainTxResponse) (*kaminarigosdk.EstimateOnChainTxResponse, error) {
	amount, err := strconv.Atoi(estimate.Fee)
	if err != nil {
		return nil, err
	}

	return &kaminarigosdk.EstimateOnChainTxResponse{
		Fee: int64(amount),
	}, nil
}

func toPbInvoice(invoice *filteredOnChainInvoice) (*kaminarigosdk.FilteredOnChainInvoice, error) {
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

	return &kaminarigosdk.FilteredOnChainInvoice{
		BitcoinAddress: invoice.BitcoinAddress,
		Description:    invoice.Description,
		Amount:         int64(amount),
		Status:         status,
		CreatedAt: &kaminarigosdk.Timestamp{
			Seconds: int64(createdAt.Second()),
		},
	}, nil
}

func toPbInvoiceStatus(status string) (kaminarigosdk.InvoiceStatus, error) {
	switch status {
	case "INVOICE_STATUS_UNSPECIFIED":
		return kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNSPECIFIED, nil
	case "INVOICE_STATUS_PAID":
		return kaminarigosdk.InvoiceStatus_INVOICE_STATUS_PAID, nil
	case "INVOICE_STATUS_UNPAID":
		return kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNPAID, nil
	}

	return 0, fmt.Errorf("invalid status")
}

func toPbTx(tx *filteredOnChainTransaction) (*kaminarigosdk.FilteredOnChainTransaction, error) {
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

	blockNumber, err := strconv.Atoi(tx.BlockNumber)
	if err != nil {
		return nil, err
	}

	return &kaminarigosdk.FilteredOnChainTransaction{
		ID:          tx.ID,
		MerchantID:  tx.MerchantID,
		Status:      status,
		Source:      tx.Source,
		Destination: tx.Destination,
		Amount:      int64(amount),
		CreatedAt: &kaminarigosdk.Timestamp{
			Seconds: int64(createdAt.Second()),
		},
		Direction:     txType,
		Confirmations: tx.Confirmations,
		BlockNumber:   int64(blockNumber),
		ExplorerUrl:   tx.ExplorerUrl,
	}, nil
}

func toPbTransactionStatus(status string) (kaminarigosdk.TransactionStatus, error) {
	switch status {
	case "TRANSACTION_STATUS_UNSPECIFIED":
		return kaminarigosdk.TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED, nil
	case "TRANSACTION_STATUS_FAILED":
		return kaminarigosdk.TransactionStatus_TRANSACTION_STATUS_FAILED, nil
	case "TRANSACTION_STATUS_SUCCESS":
		return kaminarigosdk.TransactionStatus_TRANSACTION_STATUS_SUCCESS, nil
	case "TRANSACTION_STATUS_PENDING":
		return kaminarigosdk.TransactionStatus_TRANSACTION_STATUS_PENDING, nil
	case "TRANSACTION_STATUS_WAITING_TO_FINALIZE":
		return kaminarigosdk.TransactionStatus_TRANSACTION_STATUS_WAITING_TO_FINALIZE, nil
	}

	return 0, fmt.Errorf("invalid status")
}

func toPbTransactionType(txType string) (kaminarigosdk.TransactionType, error) {
	switch txType {
	case "TRANSACTION_TYPE_UNSPECIFIED":
		return kaminarigosdk.TransactionType_TRANSACTION_TYPE_UNSPECIFIED, nil
	case "TRANSACTION_TYPE_INCOMING":
		return kaminarigosdk.TransactionType_TRANSACTION_TYPE_INCOMING, nil
	case "TRANSACTION_TYPE_OUTGOING":
		return kaminarigosdk.TransactionType_TRANSACTION_TYPE_OUTGOING, nil
	}

	return 0, fmt.Errorf("invalid type")
}
