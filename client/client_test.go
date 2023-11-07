package client_test

import (
	"encoding/base64"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
	"github.com/BoostyLabs/kaminari-go-sdk/client"
)

func TestClient(t *testing.T) {
	t.Skip("for manual testing")

	cl, err := client.DefaultClient(&client.Config{
		ApiKey:    "9fbda4b2ad024f5c98b7d21288cdcb01de83bfc9a435966cba858d6bfdf417fb",
		SecretKey: "9fbda4b2ad024f5c98b7d21288cdcb01de83bfc9a435966cba858d6bfdf417fb",
		ApiUrl:    "http://localhost:8080",
	})
	require.NoError(t, err)

	t.Run("get user balance", func(t *testing.T) {
		balance, err := cl.GetBalance(strconv.FormatInt(time.Now().Unix(), 10))
		require.NoError(t, err)
		require.EqualValues(t, balance.TotalBalance, 0)
		require.EqualValues(t, balance.FrozenAmount, 0)
	})

	t.Run("get lightning transaction", func(t *testing.T) {
		_, err := cl.GetLightningTransaction(&kaminarigosdk.GetLightningTransactionRequest{
			ID:    "95c353a5-0565-496b-83f2-5064a79268ed",
			Nonce: strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})

	t.Run("verify webhook signature", func(t *testing.T) {
		encodedEvent := "ewogICJUeXBlIjogIkVWRU5UX1RZUEVfTElHSFROSU5HX0lOVk9JQ0VfSVNfUEFJRCIsCiAgImV2ZW50X3R5cGUiOiAiRVZFTlRfVFlQRV9MSUdIVE5JTkdfSU5WT0lDRV9JU19QQUlEIiwKICAibGlnaHRuaW5nX2ludm9pY2VfaXNfcGFpZCI6IHsKICAgICJpZCI6ICI2NTIwOTg2M2RiNzhmODFlZWYyNTVjMDQ5ZmY2MzY3MzgyZmM1ZmU3ZDliMDAwYzA5ZDRkNTIwMTQ0OGQxYzRmIgogIH0KfQ=="

		event, err := base64.StdEncoding.DecodeString(encodedEvent)
		require.NoError(t, err)

		resp, err := cl.VerifyWebhookSignature(&kaminarigosdk.VerifyWebhookSignatureRequest{
			Signature: "5f0854331c3557786350f3716c8b4bc798aea73205748184ff24b642a8099ffa769dd3d42ba5db676efd8f26b5f9558b3601770349fddeea082beb4552bfebb401",
			Event:     event,
			Nonce:     strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.True(t, resp.IsValid)
	})

	t.Run("get ln invoice from lnurl", func(t *testing.T) {
		resp, err := cl.ConvertLnUrlInvoiceToLND(&kaminarigosdk.ConvertLnUrlInvoiceToLNDRequest{
			LnrulInvoice: "lnurl1dp68gurn8ghj7ampd3kx2ar0veekzar0wd5xjtnrdakj7tnhv4kxctttdehhwm30d3h82unvwqhhgmmjwp5kgung096xsmfhxvevd4tf",
			Amount:       1000,
			Nonce:        strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotEmpty(t, resp.Invoice)
	})

	t.Run("get lnurl", func(t *testing.T) {
		lnurl, err := cl.GetLightningAddress(strconv.FormatInt(time.Now().Unix(), 10))
		require.NoError(t, err)
		require.NotEmpty(t, lnurl.Invoice)
	})

	t.Run("get lnurl for merchant", func(t *testing.T) {
		lnurl, err := cl.GetLightningAddressForMerchant(&kaminarigosdk.GetLightningAddrForMerchantRequest{
			MerchantID: "Not specified",
			Nonce:      strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotEmpty(t, lnurl.Invoice)
	})

	t.Run("estimate on-chain tx", func(t *testing.T) {
		_, err := cl.EstimateOnChainTx(&kaminarigosdk.EstimateOnChainTxRequest{
			BitcoinAddress: "bcrt1q66y8c986x79gw4u86926cqw86d39m23ftacwc9",
			Amount:         100000,
			Nonce:          strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})

	var bitcoinAddress string
	t.Run("create on-chain invoice", func(t *testing.T) {
		addr, err := cl.CreateOnChainInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      100_000_000, // 1 BTC.
			Description: "test description",
			MerchantID:  "",
			Nonce:       strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotEmpty(t, addr)

		bitcoinAddress = addr
	})

	var invoice *kaminarigosdk.CreateLightningInvoiceResponse
	t.Run("create lightning invoice", func(t *testing.T) {
		resp, err := cl.CreateLightningInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      100,
			Description: "test description",
			MerchantID:  "",
			Nonce:       strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)

		invoice = resp
	})

	t.Run("get on-chain invoice", func(t *testing.T) {
		resp, err := cl.GetOnChainInvoice(&kaminarigosdk.GetOnChainInvoiceRequest{
			BitcoinAddress: bitcoinAddress,
			Nonce:          strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Invoice)
		require.NotEmpty(t, resp.Invoice.BitcoinAddress)
		require.Equal(t, "test description", resp.Invoice.Description)
		require.EqualValues(t, 100_000_000, resp.Invoice.Amount)
		require.Equal(t, kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNPAID, resp.Invoice.Status)
		require.NotEmpty(t, resp.Invoice.CreatedAt)
	})

	t.Run("get lightning invoice", func(t *testing.T) {
		resp, err := cl.GetLightningInvoice(&kaminarigosdk.GetLightningInvoiceRequest{
			ID:    invoice.ID,
			Nonce: strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp.Invoice)
		require.NotEmpty(t, resp.Invoice.ID)
		require.NotEmpty(t, resp.Invoice.EncodedInvoice)
		require.Equal(t, "test description", resp.Invoice.Description)
		require.EqualValues(t, 100, resp.Invoice.Amount)
		require.Equal(t, kaminarigosdk.InvoiceStatus_INVOICE_STATUS_UNPAID, resp.Invoice.Status)
	})

	t.Run("get on-chain transaction", func(t *testing.T) {
		_, err := cl.GetOnChainTransaction(&kaminarigosdk.GetOnChainTransactionRequest{
			ID:    "287efa8b-472f-4750-85ef-da5962889e1b",
			Nonce: strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})

	t.Run("get lightning transaction", func(t *testing.T) {
		_, err := cl.GetLightningTransaction(&kaminarigosdk.GetLightningTransactionRequest{
			ID:    "95c353a5-0565-496b-83f2-5064a79268ed",
			Nonce: strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})

	t.Run("send on-chain payment", func(t *testing.T) {
		err := cl.SendOnChainPayment(&kaminarigosdk.SendOnChainPaymentRequest{
			BitcoinAddress: "bcrt1qlpacmuy96uwzdmallnmc37wnss7n08c942mmdt",
			Amount:         700_000,
			MerchantID:     "",
			Nonce:          strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})

	t.Run("send lightning payment", func(t *testing.T) {
		err := cl.SendLightningPayment(&kaminarigosdk.SendLightningPaymentRequest{
			Invoice:    "lnbcrt5m1pj5j6hwpp5x630633cz96ufmyuz8k3e835xx6gnd3hu0lc8nacy6a4ymk4t82sdqqcqzpgxqyz5vqsp5sfam4y7yye6yp8hyk5raygjup79vd0790k6cxqvp8jk0fddgvsyq9qyyssq97x7m9yzhx9ar69vakctvex7c22ttllgmgvv6u0vfwqu28uksgtqg7qk6m0nhqrx6vzth8pzkxwyuxgn7fpr2c2apkmr98k7t2p4fucqpzxlar",
			MerchantID: "",
			Nonce:      strconv.FormatInt(time.Now().Unix(), 10),
		})
		require.NoError(t, err)
	})
}
