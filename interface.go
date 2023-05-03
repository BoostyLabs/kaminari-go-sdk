package kaminarigosdk

type Interface interface {
	CreateOnChainInvoice(*CreateInvoiceRequest) (string, error)
	CreateLightningInvoice(*CreateInvoiceRequest) (*CreateLightningInvoiceResponse, error)

	SendOnChainPayment(*SendOnChainPaymentRequest) error
	SendLightningPayment(*SendLightningPaymentRequest) error

	GetOnChainInvoice(*GetOnChainInvoiceRequest) (*GetOnChainInvoiceResponse, error)
	GetLightningInvoice(*GetLightningInvoiceRequest) (*GetLightningInvoiceResponse, error)
}

type CreateInvoiceRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	MerchantId  string `json:"merchant_id"`
}

type CreateLightningInvoiceResponse struct {
	ID      string `json:"id"`
	Invoice string `json:"invoice"`
}

type SendOnChainPaymentRequest struct {
	BitcoinAddress string `json:"bitcoin_address"`
	Amount         int64  `json:"amount"`
	MerchantId     string `json:"merchant_id"`
}

type SendLightningPaymentRequest struct {
	Invoice    string `json:"invoice"`
	MerchantId string `json:"merchant_id"`
}

type GetOnChainInvoiceRequest struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

type GetOnChainInvoiceResponse struct {
	Invoice *FilteredOnChainInvoice `json:"invoice"`
}

type FilteredOnChainInvoice struct {
	BitcoinAddress string        `json:"bitcoin_address"`
	Description    string        `json:"description"`
	Amount         int64         `json:"amount"`
	Status         InvoiceStatus `json:"status"`
	CreatedAt      *Timestamp    `json:"created_at"`
}

type InvoiceStatus int32

const (
	InvoiceStatus_INVOICE_STATUS_UNSPECIFIED InvoiceStatus = 0
	InvoiceStatus_INVOICE_STATUS_PAID        InvoiceStatus = 1
	InvoiceStatus_INVOICE_STATUS_UNPAID      InvoiceStatus = 2
)

type Timestamp struct {
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos int32 `json:"nanos,omitempty"`
}

type GetLightningInvoiceRequest struct {
	Id string `json:"id"`
}

type GetLightningInvoiceResponse struct {
	Invoice *FilteredLightningInvoice `json:"invoice"`
}

type FilteredLightningInvoice struct {
	Id             string        `json:"id"`
	EncodedInvoice string        `json:"encoded_invoice"`
	Description    string        `json:"description"`
	Amount         int64         `json:"amount"`
	Status         InvoiceStatus `json:"status"`
	CreatedAt      *Timestamp    `json:"created_at"`
}
