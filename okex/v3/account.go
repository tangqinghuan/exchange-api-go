package okex

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Accounts This endpoint supports getting the list of assets(only show pairs with balance larger than 0), the balances, amount available/on hold in spot accounts.
func (r *rest) Accounts() ([]*Account, error) {
	method := http.MethodGet
	path := "/api/spot/v3/accounts"
	content, err := r.Request(method, path, nil, nil, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "accounts request")
	}

	var as []*Account
	if err := json.Unmarshal(content, &as); err != nil {
		return nil, errors.Wrap(err, "accounts response body")
	}

	return as, nil
}

// Account This endpoint supports getting the balance, amount available/on hold of a token in spot account.
func (r *rest) Account(currency string) (*Account, error) {
	method := http.MethodGet
	path := "/api/spot/v3/accounts/" + currency
	content, err := r.Request(method, path, nil, nil, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "account request")
	}

	var a *Account
	if err := json.Unmarshal(content, &a); err != nil {
		return nil, errors.Wrap(err, "account response body")
	}

	return a, nil
}
