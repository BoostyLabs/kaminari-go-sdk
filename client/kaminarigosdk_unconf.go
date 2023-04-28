package client

import (
	"errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

var (
	ErrNotConfigured = errors.New("kaminari not configured")
)

type hNotConf struct{}

// NewNotConf creates not configured kaminari object which returns an error on each response.
func NewNotConf() kaminarigosdk.Interface {
	return &hNotConf{}
}

func (h *hNotConf) CreateOnChainInvoice(req *kaminarigosdk.CreateInvoiceRequest) (string, error) {
	return "", ErrNotConfigured
}

func (h *hNotConf) CreateLightningInvoice(req *kaminarigosdk.CreateInvoiceRequest) (*kaminarigosdk.CreateLightningInvoiceResponse, error) {
	return nil, ErrNotConfigured
}
