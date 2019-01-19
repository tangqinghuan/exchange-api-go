package okex

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// AllTicker Get the last traded price, best bid/ask price, 24 hour trading volume and more info of all trading pairs.
func (r *rest) AllTicker() ([]*Ticker, error) {
	method := http.MethodGet
	path := "/api/spot/v3/instruments/ticker"
	content, err := r.Request(method, path, nil, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "all ticker request")
	}

	var ts []*Ticker
	if err := json.Unmarshal(content, &ts); err != nil {
		return nil, errors.Wrap(err, "all ticker response body")
	}

	return ts, nil
}

// Ticker Get the last traded price, best bid/ask price, 24 hour trading volume and more info of a trading pair.
func (r *rest) Ticker(instrumentID string) (*Ticker, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/api/spot/v3/instruments/%s/ticker", instrumentID)
	content, err := r.Request(method, path, nil, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "ticker request")
	}

	var tk *Ticker
	if err := json.Unmarshal(content, &tk); err != nil {
		return nil, errors.Wrap(err, "ticker response body")
	}

	return tk, nil
}
