package okex

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// Instruments Get market data. This endpoint provides the snapshots of market data and can be used without verifications.
// List trading pairs and get the trading limit, price, and more information of different trading pairs.
// GET /api/spot/v3/instruments
func (r *rest) Instruments() ([]*Instrument, error) {
	method := http.MethodGet
	path := "/api/spot/v3/instruments"
	content, err := r.Request(method, path, nil, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "instruments request")
	}

	var ss []*Instrument
	if err := json.Unmarshal(content, &ss); err != nil {
		return nil, errors.Wrap(err, "Instruments response body")
	}

	return ss, nil
}
