package client

import (
	"fmt"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (c *Client) verifyWebhookSignature(req *kaminarigosdk.VerifyWebhookSignatureRequest) (*kaminarigosdk.VerifyWebhookSignatureResponse, error) {
	url := fmt.Sprintf("%s/gateway/api/webhook/verification", c.cfg.ApiUrl)
	var result kaminarigosdk.VerifyWebhookSignatureResponse

	resp, err := c.restyClient.R().
		SetBody(req.Event).
		SetHeader("Kaminari-Signature", req.Signature).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, err
	}

	return &result, nil
}
