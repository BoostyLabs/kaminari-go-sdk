package client_test

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
	"github.com/BoostyLabs/kaminari-go-sdk/client"
)

func TestClient(t *testing.T) {
	t.Skip("for manual testing")

	cl, err := client.DefaultClient(&client.Config{
		ApiKey: "9fbda4b2ad024f5c98b7d21288cdcb01de83bfc9a435966cba858d6bfdf417fb",
		ApiUrl: "http://localhost:8080",
	})
	require.NoError(t, err)

	t.Run("get user balance", func(t *testing.T) {
		balance, err := cl.GetBalance()
		require.NoError(t, err)
		require.EqualValues(t, balance.TotalBalance, 0)
		require.EqualValues(t, balance.FrozenAmount, 0)
	})

	t.Run("get lnurl", func(t *testing.T) {
		lnurl, err := cl.GetLightningAddress()
		require.NoError(t, err)
		require.NotEmpty(t, lnurl.Invoice)
	})

	t.Run("get lnurl for merchant", func(t *testing.T) {
		lnurl, err := cl.GetLightningAddressForMerchant(&kaminarigosdk.GetLightningAddrForMerchantRequest{
			MerchantID: "Not specified",
		})
		require.NoError(t, err)
		require.NotEmpty(t, lnurl.Invoice)
	})

	t.Run("estimate on-chain tx", func(t *testing.T) {
		_, err := cl.EstimateIOChainTx(&kaminarigosdk.EstimateOnChainTxRequest{
			BitcoinAddress: "bcrt1q66y8c986x79gw4u86926cqw86d39m23ftacwc9",
			Amount:         10000,
		})
		require.NoError(t, err)
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
		encodedEvent := "ewoJIlR5cGUiOiAiRVZFTlRfVFlQRV9MSUdIVE5JTkdfSU5WT0lDRV9JU19QQUlEIiwKCSJsaWdodG5pbmdfaW52b2ljZV9pc19wYWlkIjogewoJCSJpZCI6ICJ0ZXN0X2lkIgoJfQp9"

		event, err := base64.StdEncoding.DecodeString(encodedEvent)
		require.NoError(t, err)

		resp, err := cl.VerifyWebhookSignature(&kaminarigosdk.VerifyWebhookSignatureRequest{
			Signature: "8822c5e52859e6850381749975b8eacb3b980c8a6b668abbc89a9a7117a0754e3f7778d97cc1e83375686b041c4a1f0dd8b901265f983ec9e56f1a38c53450fe01",
			Event:     event,
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.True(t, resp.IsValid)
	})
}
