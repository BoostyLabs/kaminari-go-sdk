package client_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
	"github.com/BoostyLabs/kaminari-go-sdk/client"
)

func TestClient(t *testing.T) {
	t.Skip("for manual testing")

	cl := client.DefaultClient(&client.Config{
		ApiKey: "8a3093744c910eb0dad59d52556bcfc35036e32a8868019a05c8352cad226ebe",
		ApiUrl: "http://localhost:8080",
	})

	var bitcoinAddress string
	t.Run("create on-chain invoice", func(t *testing.T) {
		addr, err := cl.CreateOnChainInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantID:  "",
		})
		require.NoError(t, err)
		require.NotEmpty(t, addr)

		bitcoinAddress = addr
	})

	var invoice *kaminarigosdk.CreateLightningInvoiceResponse
	t.Run("create lightning invoice", func(t *testing.T) {
		resp, err := cl.CreateLightningInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantID:  "",
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp)

		invoice = resp
	})

	t.Run("send on-chain payment", func(t *testing.T) {
		err := cl.SendOnChainPayment(&kaminarigosdk.SendOnChainPaymentRequest{
			BitcoinAddress: bitcoinAddress,
			Amount:         1,
			MerchantID:     "",
		})
		require.NoError(t, err)
	})

	t.Run("send lightning payment", func(t *testing.T) {
		err := cl.SendLightningPayment(&kaminarigosdk.SendLightningPaymentRequest{
			Invoice:    "lnbcrt30n1pj92rhdpp59222w4qnt7j7q9e5q8h4ccf6hpuhmuw3hxj8rtta3galm0g7jrxqdq6w3jhxazlv3jhxcmjd9c8g6t0dccqzpgxqyz5vqsp5342yghgpm639hywl90a6xqmkjn2xyf2xewmn6lk27u7nma5zspzs9qyyssqv77n7dvyu7m4h54eqmuq69z4nc2299tnldj0qrddmj6mmzfpcdvq37dahw6kandj9c4turp2j8rgflyyfruuagkl5truug2977uwkdspsej75a",
			MerchantID: "",
		})
		require.NoError(t, err)
	})

	t.Run("get on-chain invoice", func(t *testing.T) {
		resp, err := cl.GetOnChainInvoice(&kaminarigosdk.GetOnChainInvoiceRequest{
			BitcoinAddress: bitcoinAddress,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Invoice)
		require.NotEmpty(t, resp.Invoice.BitcoinAddress)
		require.Equal(t, "test description", resp.Invoice.Description)
		require.EqualValues(t, 1, resp.Invoice.Amount)
		require.Equal(t, kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNPAID, resp.Invoice.Status)
		require.NotEmpty(t, resp.Invoice.CreatedAt)
	})

	t.Run("get lightning invoice", func(t *testing.T) {
		resp, err := cl.GetLightningInvoice(&kaminarigosdk.GetLightningInvoiceRequest{
			ID: invoice.ID,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Invoice)
		require.NotEmpty(t, resp.Invoice.ID)
		require.NotEmpty(t, resp.Invoice.EncodedInvoice)
		require.Equal(t, "test description", resp.Invoice.Description)
		require.EqualValues(t, 1, resp.Invoice.Amount)
		require.Equal(t, kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNPAID, resp.Invoice.Status)
		require.NotEmpty(t, resp.Invoice.CreatedAt)
	})

	t.Run("get on-chain transaction", func(t *testing.T) {
		_, err := cl.GetOnChainTransaction(&kaminarigosdk.GetOnChainTransactionRequest{
			ID: "684cc668-96de-416c-9ff9-22531f5b6899",
		})
		require.NoError(t, err)
	})

	t.Run("get lightning transaction", func(t *testing.T) {
		_, err := cl.GetLightningTransaction(&kaminarigosdk.GetLightningTransactionRequest{
			ID: "3530ab8c-b721-4a15-97b8-0cada5546b3b",
		})
		require.NoError(t, err)
	})

	t.Run("verify webhook signature", func(t *testing.T) {
		resp, err := cl.VerifyWebhookSignature(&kaminarigosdk.VerifyWebhookSignatureRequest{
			Signature: "",
			Event:     &kaminarigosdk.Event{},
		})
		require.Error(t, err)
		require.Nil(t, resp)
	})
}