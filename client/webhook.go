package client

import (
	"fmt"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

func (c *Client) VerifyWebhookSignature(req *kaminarigosdk.VerifyWebhookSignatureRequest) (*kaminarigosdk.VerifyWebhookSignatureResponse, error) {
	url := fmt.Sprintf("%s/api/webhook/signature/verify", c.cfg.ApiUrl)
	var result kaminarigosdk.VerifyWebhookSignatureResponse

	resp, err := c.restyClient.R().
		SetBody(req).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't verify webhook signature")
	}

	return &result, nil
}
