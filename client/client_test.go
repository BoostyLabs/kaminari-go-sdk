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
		ApiKey: "9d326b15f6923007ab8138237a646b1f48f620f32179801ab334ee1026918a89",
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
		require.Error(t, err)
	})

	t.Run("send lightning payment", func(t *testing.T) {
		err := cl.SendLightningPayment(&kaminarigosdk.SendLightningPaymentRequest{
			Invoice:    invoice.Invoice,
			MerchantID: "",
		})
		require.Error(t, err)
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
		require.Equal(t, "1", resp.Invoice.Amount)
		require.Equal(t, "INVOICE_STATUS_UNPAID", resp.Invoice.Status)
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
		require.Equal(t, "1", resp.Invoice.Amount)
		require.Equal(t, "INVOICE_STATUS_UNPAID", resp.Invoice.Status)
		require.NotEmpty(t, resp.Invoice.CreatedAt)
	})

	t.Run("get on-chain transaction", func(t *testing.T) {
		_, err := cl.GetOnChainTransaction(&kaminarigosdk.GetOnChainTransactionRequest{
			ID: "",
		})
		require.Error(t, err)
	})

	t.Run("get lightning transaction", func(t *testing.T) {
		_, err := cl.GetLightningTransaction(&kaminarigosdk.GetLightningTransactionRequest{
			ID: "",
		})
		require.Error(t, err)
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
