package okex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// NewOrder OKEx token trading only supports limit and market orders (more order types will become available in the future). You can place an order only if you have enough funds.
// Once your order is placed, the amount will be put on hold.
func (r *rest) NewOrder(req *OrderRequest) (*OrderResponse, error) {
	method := http.MethodPost
	path := "/api/spot/v3/orders"
	content, err := r.Request(method, path, nil, req, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "new order request")
	}

	var res *OrderResponse
	if err := json.Unmarshal(content, &res); err != nil {
		return nil, errors.Wrap(err, "new order response body")
	}

	return res, nil
}

// BatchNewOrder This endpoint supports placing multiple orders for specific trading pairs( up to 4 trading pairs, maximum 4 orders for each pair).
func (r *rest) BatchNewOrder(req []*OrderRequest) (map[string][]*OrderResponse, error) {
	method := http.MethodPost
	path := "/api/spot/v3/batch_orders"
	content, err := r.Request(method, path, nil, req, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "batch new order request")
	}

	var res map[string][]*OrderResponse
	if err := json.Unmarshal(content, &res); err != nil {
		return nil, errors.Wrap(err, "batch new order response body")
	}

	return res, nil
}

// 2018-10-12T07:32:56.512ZPOST/api/spot/v3/cancel_orders/1611729012263936{"client_oid":"20181009","instrument_id":"btc-usdt"}

// CancelOrder Cancelling an unfilled order.
// instrument_id [required]By providing this parameter, the corresponding order of a designated trading pair will be cancelled. If not providing this parameter, it will be back to error code.
// client_oid [optional]the order ID created by yourself
// order_id [required]order ID
func (r *rest) CancelOrder(instrumentID, clientOID, orderID string) (*OrderResponse, error) {
	method := http.MethodPost
	path := "/api/spot/v3/cancel_orders/" + orderID
	data := make(map[string]string)
	data["instrument_id"] = instrumentID
	data["client_oid"] = clientOID
	content, err := r.Request(method, path, nil, data, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "cancel order request")
	}

	var res *OrderResponse
	if err := json.Unmarshal(content, &res); err != nil {
		return nil, errors.Wrap(err, "cancel order response body")
	}

	return res, nil
}

// BatchCancelOrder With best effort, this endpoints supports cancelling all open orders for a specific trading pair or several trading pairs.
func (r *rest) BatchCancelOrder(req []*BatchCancelOrderRequest) (map[string]*BatchCancelOrderResponse, error) {
	method := http.MethodPost
	path := "/api/spot/v3/cancel_batch_orders"
	content, err := r.Request(method, path, nil, req, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "batch cancel order request")
	}

	var res map[string]*BatchCancelOrderResponse
	if err := json.Unmarshal(content, &res); err != nil {
		return nil, errors.Wrap(err, "batch cancel order response body")
	}

	return res, nil
}

// OrderHistory List your orders. Cursor pagination is used. All paginated requests return the latest information (newest) as the first page sorted by newest (in chronological time) first.
// status [required] list the status of all orders (all, open, part_filled, canceling, filled, cancelled，ordering,failure)
// instrument_id [required] list the orders of specific trading pairs
// from [optional]request page after this id (latest information) (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
// to [optional]request page after (older) this id.
// limit [optional]number of results per request. Maximum 100. (default 100)
func (r *rest) OrderHistory(instrumentID, fromID, toID string, limit int32, status []string) ([]*Order, error) {
	method := http.MethodGet
	path := "/api/spot/v3/orders"
	params := make(map[string]string)
	params["instrument_id"] = instrumentID
	params["from"] = fromID
	params["to"] = toID
	if limit != 0 {
		params["limit"] = fmt.Sprint(limit)
	}
	params["status"] = strings.Join(status, "|")
	content, err := r.Request(method, path, params, nil, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "order history request")
	}

	var ods []*Order
	if err := json.Unmarshal(content, &ods); err != nil {
		return nil, errors.Wrap(err, "order history response body")
	}

	return ods, nil
}

// OrderPending List all your current open orders. Cursor pagination is used. All paginated requests return the latest information (newest) as the first page sorted by newest (in chronological time) first.
// from [optional]request page after this id (latest information) (eg. 1, 2, 3, 4, 5. There is only a 5 "from 4", while there are 1, 2, 3 "to 4")
// to [optional]request page after (older) this id.
// limit [optional]number of results per request. Maximum 100. (default 100)
// instrument_id [optional]trading pair ,information of all trading pair will be returned if the field is left blank
func (r *rest) OrderPending(instrumentID, fromID, toID string, limit int32) ([]*Order, error) {
	method := http.MethodGet
	path := "/api/spot/v3/orders_pending"
	params := make(map[string]string)
	params["instrument_id"] = instrumentID
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
		return nil, errors.Wrap(err, "order pending request")
	}

	var ods []*Order
	if err := json.Unmarshal(content, &ods); err != nil {
		return nil, errors.Wrap(err, "order pending response body")
	}

	return ods, nil
}
