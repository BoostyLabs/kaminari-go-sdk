package client

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"

	kaminarigosdk "github.com/BoostyLabs/kaminari-go-sdk"
)

type balance struct {
	TotalBalance string `json:"totalBalance"`
	FrozenAmount string `json:"frozenAmount"`
}

func (c *Client) GetBalance() (*kaminarigosdk.Balance, error) {
	url := fmt.Sprintf("%s/api/lightning/v1/balance", c.cfg.ApiUrl)
	var balanceResp balance
	resp, err := c.restyClient.R().
		SetResult(&balanceResp).
		Get(url)
	if err := checkForError(resp, err); err != nil {
		return nil, errors.Wrap(err, "can't get user balance")
	}

	return toPbBalance(&balanceResp)
}

func toPbBalance(balance *balance) (*kaminarigosdk.Balance, error) {
	totalBalance, err := strconv.Atoi(balance.TotalBalance)
	if err != nil {
		return nil, err
	}

	frozenBalance, err := strconv.Atoi(balance.TotalBalance)
	if err != nil {
		return nil, err
	}

	return &kaminarigosdk.Balance{
		TotalBalance: int64(totalBalance),
		FrozenAmount: int64(frozenBalance),
	}, nil
}
