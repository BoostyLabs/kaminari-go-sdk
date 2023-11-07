package client

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

const ApiSignatureHeader = "X-kaminari-api-signature"

type Client struct {
	cfg         *Config
	restyClient *resty.Client
}

func DefaultClient(cfg *Config) (kaminarigosdk.Interface, error) {
	if !isCfgValid(cfg) {
		err := errors.Errorf("kaminari config is not valid")
		return nil, err
	}

	restyClient := resty.New().
		SetRetryCount(3).
		SetHeader("X-kaminari-api-key", cfg.ApiKey).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(15 * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				return r != nil && r.StatusCode() == http.StatusTooManyRequests
			},
		)

	return &Client{
		cfg:         cfg,
		restyClient: restyClient,
	}, nil
}

// GetSignature produces signature for API authorization.
func (client *Client) GetSignature(rawUrl string, nonce string, body []byte) (string, error) {
	// how to ge to signature
	// hex(secretKey(url + sha256(nonce + body)))

	parsedURL, err := url.Parse(rawUrl)
	if err != nil {
		return "", nil
	}

	// nonce + body => sha256
	sha := sha256.New()
	sha.Write([]byte(nonce))
	sha.Write(body)
	sum := sha.Sum(nil)

	secretKeyBytes, err := hex.DecodeString(client.cfg.SecretKey)
	if err != nil {
		return "", err
	}

	// append url to hmac all params url + sha256(nonce, body), and signs them.
	mac := hmac.New(sha512.New, secretKeyBytes)
	mac.Write(append([]byte(parsedURL.String()), sum...))
	macSum := mac.Sum(nil)

	signature := hex.EncodeToString(macSum)

	return signature, nil
}
