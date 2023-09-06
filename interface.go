package kaminarigosdk

type Interface interface {
	GetBalance() (*Balance, error)

	EstimateIOChainTx(req *EstimateOnChainTxRequest) (*EstimateOnChainTxResponse, error)

	CreateOnChainInvoice(*CreateInvoiceRequest) (string, error)
	CreateLightningInvoice(*CreateInvoiceRequest) (*CreateLightningInvoiceResponse, error)

	SendOnChainPayment(*SendOnChainPaymentRequest) error
	SendLightningPayment(*SendLightningPaymentRequest) error

	GetOnChainInvoice(*GetOnChainInvoiceRequest) (*GetOnChainInvoiceResponse, error)
	GetLightningInvoice(*GetLightningInvoiceRequest) (*GetLightningInvoiceResponse, error)

	GetOnChainTransaction(*GetOnChainTransactionRequest) (*GetOnChainTransactionResponse, error)
	GetLightningTransaction(*GetLightningTransactionRequest) (*GetLightningTransactionResponse, error)

	VerifyWebhookSignature(*VerifyWebhookSignatureRequest) (*VerifyWebhookSignatureResponse, error)
}

type CreateInvoiceRequest struct {
	Amount      int64  `json:"amount"`
	Description string `json:"description"`
	MerchantID  string `json:"merchant_id"`
}

type CreateLightningInvoiceResponse struct {
	ID      string `json:"id"`
	Invoice string `json:"invoice"`
}

type SendOnChainPaymentRequest struct {
	BitcoinAddress string `json:"bitcoin_address"`
	Amount         int64  `json:"amount"`
	MerchantID     string `json:"merchant_id"`
}

type SendLightningPaymentRequest struct {
	Invoice    string `json:"invoice"`
	MerchantID string `json:"merchant_id"`
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
	ID string `json:"id"`
}

type GetLightningInvoiceResponse struct {
	Invoice *FilteredLightningInvoice `json:"invoice"`
}

type FilteredLightningInvoice struct {
	ID             string        `json:"id"`
	EncodedInvoice string        `json:"encoded_invoice"`
	Description    string        `json:"description"`
	Amount         int64         `json:"amount"`
	Status         InvoiceStatus `json:"status"`
	CreatedAt      *Timestamp    `json:"created_at"`
}

type GetOnChainTransactionRequest struct {
	ID string `json:"id"`
}

type GetOnChainTransactionResponse struct {
	Transaction *FilteredOnChainTransaction `json:"transaction"`
}

type FilteredOnChainTransaction struct {
	ID            string            `json:"id"`
	MerchantID    string            `json:"merchant_id"`
	Status        TransactionStatus `json:"status"`
	Source        string            `json:"source"`
	Destination   string            `json:"destination"`
	Amount        int64             `json:"amount"`
	CreatedAt     *Timestamp        `json:"created_at"`
	Direction     TransactionType   `json:"direction"`
	Confirmations int32             `json:"confirmations"`
	BlockNumber   int64             `json:"block_number"`
	ExplorerUrl   string            `json:"explorer_url"`
}

type TransactionStatus int32

const (
	TransactionStatus_TRANSACTION_STATUS_UNSPECIFIED         TransactionStatus = 0
	TransactionStatus_TRANSACTION_STATUS_FAILED              TransactionStatus = 1
	TransactionStatus_TRANSACTION_STATUS_SUCCESS             TransactionStatus = 2
	TransactionStatus_TRANSACTION_STATUS_PENDING             TransactionStatus = 3
	TransactionStatus_TRANSACTION_STATUS_WAITING_TO_FINALIZE TransactionStatus = 4
)

type TransactionType int32

const (
	TransactionType_TRANSACTION_TYPE_UNSPECIFIED TransactionType = 0
	TransactionType_TRANSACTION_TYPE_INCOMING    TransactionType = 1
	TransactionType_TRANSACTION_TYPE_OUTGOING    TransactionType = 2
)

type GetLightningTransactionRequest struct {
	ID string `json:"id"`
}

type GetLightningTransactionResponse struct {
	Transaction *FilteredLightningTransaction `json:"transaction"`
}

type FilteredLightningTransaction struct {
	ID          string            `json:"id"`
	MerchantID  string            `json:"merchant_id"`
	Status      TransactionStatus `json:"status"`
	Source      string            `json:"source"`
	Destination string            `json:"destination"`
	Amount      int64             `json:"amount"`
	CreatedAt   *Timestamp        `json:"created_at"`
	Direction   TransactionType   `json:"direction"`
	ExplorerUrl string            `json:"explorer_url"`
}

type VerifyWebhookSignatureRequest struct {
	Signature string `json:"signature"`
	Event     []byte `json:"event"`
}

type Event struct {
	EventType              EventType               `json:"event_type"`
	LightningInvoiceIsPaid *LightningInvoiceIsPaid `json:"lightning_invoice_is_paid,omitempty"`
	BitcoinInvoiceIsPaid   *BitcoinInvoiceIsPaid   `json:"bitcoin_invoice_is_paid,omitempty"`
}

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED               EventType = 0
	EventType_EVENT_TYPE_LIGHTNING_INVOICE_IS_PAID EventType = 1
	EventType_EVENT_TYPE_BITCOIN_INVOICE_IS_PAID   EventType = 2
)

type LightningInvoiceIsPaid struct {
	Id string `json:"id"`
}

type BitcoinInvoiceIsPaid struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

type VerifyWebhookSignatureResponse struct {
	IsValid bool `json:"is_valid"`
}

type Balance struct {
	TotalBalance int64 `json:"totalBalance"`
	FrozenAmount int64 `json:"frozenAmount"`
}

type EstimateOnChainTxRequest struct {
	BitcoinAddress string `json:"bitcoin_address"`
	Amount         int64  `json:"amount"`
}

type EstimateOnChainTxResponse struct {
	Fee int64 `json:"amount"`
}
