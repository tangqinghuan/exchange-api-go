package okex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Accounts This endpoint supports getting the list of assets(only show pairs with balance larger than 0), the balances, amount available/on hold in spot accounts.
// GET /api/spot/v3/accounts
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
// GET /api/spot/v3/accounts/<currency>
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

// Bills All paginated requests return the latest information (newest) as the first page sorted by newest (in chronological time) first.
// from [optional]request page before(newer) this id.
// to [optional]request page after(older) this id.
// limit [optional]number of results per request. Maximum 100.(default 100)
// GET /api/spot/v3/accounts/<currency>/ledger
func (r *rest) Bills(currency, fromID, toID string, limit int) ([]*Bill, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/api/spot/v3/accounts/%s/ledger", currency)
	params := make(map[string]string)
	params["from"] = fromID
	params["to"] = toID
	if limit != 0 {
		params["limit"] = fmt.Sprint(limit)
	}
	content, err := r.Request(method, path, params, nil, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "bills request")
	}

	var bs []*Bill
	if err := json.Unmarshal(content, &bs); err != nil {
		return nil, errors.Wrap(err, "bills response body")
	}

	return bs, nil
}
