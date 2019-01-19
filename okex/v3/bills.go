package okex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Bills All paginated requests return the latest information (newest) as the first page sorted by newest (in chronological time) first.
// from [optional]request page before(newer) this id.
// to [optional]request page after(older) this id.
// limit [optional]number of results per request. Maximum 100.(default 100)
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
