package okex

import (
	"encoding/json"
	"net/http"

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
// instrument_id	string	Yes	By providing this parameter, the corresponding order of a designated trading pair will be cancelled. If not providing this parameter, it will be back to error code.
// client_oid	string	No	the order ID created by yourself
// order_id	string	Yes	order ID
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
func (r *rest) BatchCancelOrder(req []*BatchCancelOrderRequest) ([]*BatchCancelOrderResponse, error) {
	method := http.MethodPost
	path := "/api/spot/v3/cancel_batch_orders"
	content, err := r.Request(method, path, nil, req, true)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "batch cancel order request")
	}

	var res []*BatchCancelOrderResponse
	if err := json.Unmarshal(content, &res); err != nil {
		return nil, errors.Wrap(err, "batch cancel order response body")
	}

	return res, nil
}
