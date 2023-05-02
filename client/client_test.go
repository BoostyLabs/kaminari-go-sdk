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
		ApiKey: "9fbda4b2ad024f5c98b7d21288cdcb01de83bfc9a435966cba858d6bfdf417fb",
		ApiUrl: "http://localhost:8080",
	})

	var bitcoinAddress string
	t.Run("create on-chain invoice", func(t *testing.T) {
		addr, err := cl.CreateOnChainInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantId:  "",
		})
		require.NoError(t, err)
		require.NotEmpty(t, addr)

		bitcoinAddress = addr
	})

	var lightningInvoice string
	t.Run("create lightning invoice", func(t *testing.T) {
		resp, err := cl.CreateLightningInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantId:  "",
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp)

		lightningInvoice = resp.Invoice
	})

	t.Run("send on-chain payment", func(t *testing.T) {
		err := cl.SendOnChainPayment(&kaminarigosdk.SendOnChainPaymentRequest{
			BitcoinAddress: bitcoinAddress,
			Amount:         1,
			MerchantId:     "",
		})
		require.NoError(t, err)
	})

	t.Run("send lightning payment", func(t *testing.T) {
		err := cl.SendLightningPayment(&kaminarigosdk.SendLightningPaymentRequest{
			Invoice:    lightningInvoice,
			MerchantId: "",
		})
		require.NoError(t, err)
	})
}
