package okex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Trades Get the recent 60 transactions of all trading pairs. Cursor pagination is used.
// All paginated requests return the latest information (newest) as the first page sorted by newest (in chronological time) first.
// from [optional]request page after this id (latest information)
// (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
// to [optional]request page after (older) this id.
// limit [optional]number of results per request. Maximum 100. (default 100)
// instrument_id [required] trading pairs
// GET /api/spot/v3/instruments/<instrument_id>/trades
func (r *rest) Trades(instrumentID, fromID, toID string, limit int) ([]*Trade, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/api/spot/v3/instruments/%s/trades", instrumentID)
	params := make(map[string]string)
	params["instrument_id"] = instrumentID
	params["from"] = fromID
	params["to"] = toID
	if limit != 0 {
		params["limit"] = fmt.Sprint(limit)
	}
	content, err := r.Request(method, path, params, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "trades request")
	}

	var ts []*Trade
	if err := json.Unmarshal(content, &ts); err != nil {
		return nil, errors.Wrap(err, "trades response body")
	}

	return ts, nil
}
