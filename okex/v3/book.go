package okex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// OrderBook Getting the order book of a trading pair. Pagination is not supported here. The whole book will be returned for one request. WebSocket is recommended here.
// size [optional]number of results per request. Maximum 200
// depth [optional]the aggregation of the book. e.g . 0.1,0.001
// instrument_id [required] trading pairs
func (r *rest) OrderBook(instrumentID, depth string, size int) (*OrderBook, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/api/spot/v3/instruments/%s/book", instrumentID)
	params := make(map[string]string)
	params["depth"] = depth
	if size != 0 {
		params["size"] = fmt.Sprint(size)
	}
	content, err := r.Request(method, path, params, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "order book request")
	}

	var ob *OrderBook
	if err := json.Unmarshal(content, &ob); err != nil {
		return nil, errors.Wrap(err, "order book response body")
	}

	return ob, nil
}
