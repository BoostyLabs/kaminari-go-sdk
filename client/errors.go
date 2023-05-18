package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type kaminariError struct {
	GrpcCode string `json:"code"`
	Message  string `json:"message"`
}

func parseKaminariError(rawError []byte) (*kaminariError, error) {
	var err kaminariError
	if err := json.Unmarshal(rawError, &err); err != nil {
		return nil, errors.Wrap(err, "can't unmarshal raw error")
	}

	return &err, nil
}

func (err *kaminariError) Error() string {
	tmpl := `
Code:        %v
Message:       %v
`
	return fmt.Sprintf(tmpl, err.GrpcCode, err.Message)
}

type httpErr struct {
	HttpStatus      int
	UnderlyingError error
}

func newHttpErr(status int, underlyingError error) error {
	return &httpErr{
		HttpStatus:      status,
		UnderlyingError: underlyingError,
	}
}

func (err *httpErr) Error() string {
	tmpl := `
Status:          %v
UnderlyingError: %v
	`
	return fmt.Sprintf(tmpl, err.HttpStatus, err.UnderlyingError)
}

func checkForError(resp *resty.Response, err error) error {
	if err != nil {
		return err
	}

	if !isSuccess(resp.StatusCode()) {
		kaminariErr, err := parseKaminariError(resp.Body())
		if err != nil {
			return newHttpErr(resp.StatusCode(), errors.New(resp.String()))
		}

		return newHttpErr(resp.StatusCode(), kaminariErr)
	}

	return nil
}

func isSuccess(code int) bool {
	return code >= http.StatusOK && code < 300
}
