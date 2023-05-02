package kaminarigosdk

type Interface interface {
	CreateOnChainInvoice(*CreateInvoiceRequest) (string, error)
	CreateLightningInvoice(*CreateInvoiceRequest) (*CreateLightningInvoiceResponse, error)
}

type CreateInvoiceRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	MerchantId  string `json:"merchant_id"`
}

type CreateLightningInvoiceResponse struct {
	ID      int64  `json:"id"`
	Invoice string `json:"invoice"`
}
