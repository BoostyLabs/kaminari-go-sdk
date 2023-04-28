//go:build !mock && !scale
// +build !mock,!scale

package client_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
	"github.com/BoostyLabs/kaminari-go-sdk/client"
)

func TestAdapter(t *testing.T) {
	t.Skip("for manual testing")

	cl := client.DefaultAdapter(&client.Config{
		ApiKey: "9fbda4b2ad024f5c98b7d21288cdcb01de83bfc9a435966cba858d6bfdf417fb",
		ApiUrl: "http://localhost:8080",
	})

	t.Run("create on-chain invoice", func(t *testing.T) {
		addr, err := cl.CreateOnChainInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantId:  "",
		})
		require.NoError(t, err)
		require.NotEmpty(t, addr)
	})

	t.Run("create lightning invoice", func(t *testing.T) {
		resp, err := cl.CreateLightningInvoice(&kaminarigosdk.CreateInvoiceRequest{
			Amount:      1,
			Description: "test description",
			MerchantId:  "",
		})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.NotNil(t, resp)
	})
}
