package kaminarigosdk

type Interface interface {
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
	BitcoinAddress string `json:"bitcoinAddress"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}

type GetLightningInvoiceRequest struct {
	ID string `json:"id"`
}

type GetLightningInvoiceResponse struct {
	Invoice *FilteredLightningInvoice `json:"invoice"`
}

type FilteredLightningInvoice struct {
	ID             string `json:"id"`
	EncodedInvoice string `json:"encodedInvoice"`
	Description    string `json:"description"`
	Amount         string `json:"amount"`
	Status         string `json:"status"`
	CreatedAt      string `json:"createdAt"`
}

type GetOnChainTransactionRequest struct {
	ID string `json:"id"`
}

type GetOnChainTransactionResponse struct {
	Transaction *FilteredOnChainTransaction `json:"transaction"`
}

type FilteredOnChainTransaction struct {
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

type GetLightningTransactionRequest struct {
	ID string `json:"id"`
}

type GetLightningTransactionResponse struct {
	Transaction *FilteredLightningTransaction `json:"transaction"`
}

type FilteredLightningTransaction struct {
	ID          string `json:"id"`
	MerchantID  string `json:"merchantId"`
	Status      string `json:"status"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Amount      string `json:"amount"`
	CreatedAt   string `json:"createdAt"`
	Direction   string `json:"direction"`
	ExplorerUrl string `json:"explorerUrl"`
}

type VerifyWebhookSignatureRequest struct {
	Signature string `json:"signature"`
	Event     *Event `json:"event"`
}

type Event struct {
	EventType    EventType     `json:"event_type"`
	EventPayload *EventPayload `json:"event_payload"`
}

type EventType int32

const (
	EventType_EVENT_TYPE_UNSPECIFIED               EventType = 0
	EventType_EVENT_TYPE_LIGHTNING_INVOICE_IS_PAID EventType = 1
	EventType_EVENT_TYPE_BITCOIN_INVOICE_IS_PAID   EventType = 2
)

type EventPayload struct {
	LightningInvoiceIsPaid *LightningInvoiceIsPaid `json:"lightning_invoice_is_paid"`
	BitcoinInvoiceIsPaid   *BitcoinInvoiceIsPaid   `json:"bitcoin_invoice_is_paid"`
}

type LightningInvoiceIsPaid struct {
	Id string `json:"id"`
}

type BitcoinInvoiceIsPaid struct {
	BitcoinAddress string `json:"bitcoin_address"`
}

type VerifyWebhookSignatureResponse struct {
	IsValid bool `json:"is_valid"`
}
