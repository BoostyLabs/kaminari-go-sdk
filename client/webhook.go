package client

import (
	"encoding/json"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type verifyWebhookSignatureResponse struct {
	IsValid bool `json:"isValid"`
}

func (client *Client) VerifyWebhookSignature(req *kaminarigosdk.VerifyWebhookSignatureRequest) (*kaminarigosdk.VerifyWebhookSignatureResponse, error) {
	uriPath := "/api/webhooks/v1/signature/verify"
	url := client.cfg.ApiUrl + uriPath

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	signature, err := client.GetSignature(uriPath, req.Nonce, body)
	if err != nil {
		return nil, err
	}

	var result verifyWebhookSignatureResponse

	resp, err := client.restyClient.R().
		SetBody(req).
		SetHeader(ApiSignatureHeader, signature).
		SetResult(&result).
		Post(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't verify webhook signature")
	}

	return &kaminarigosdk.VerifyWebhookSignatureResponse{
		IsValid: result.IsValid,
	}, nil
}
