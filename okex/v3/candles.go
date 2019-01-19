package okex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Candles Get the charts of the trading pairs. Charts are returned in grouped buckets based on requested granularity.
// The granularity field must be one of the following values: {60 180 300 900 1800 3600 7200 14400 21600 43200 86400 604800}
// Otherwise, your request will be rejected.
// These values correspond to timeslices representing one minute, three minutes, five minutes, fifteen minutes, thirty minutes, one hour, two hours, six hours, twelve hours, one day, and 1 week respectively.
// start time in ISO 8601
// end time in ISO 8601
func (r *rest) Candles(instrumentID string, granularity int32, start, end *time.Time) ([]*Candle, error) {
	method := http.MethodGet
	path := fmt.Sprintf("/api/spot/v3/instruments/%s/candles", instrumentID)
	params := make(map[string]string)
	params["granularity"] = fmt.Sprint(granularity)
	if start != nil {
		params["start"] = start.UTC().Format("2006-01-02T15:04:05.999Z07:00")
	}
	if end != nil {
		params["end"] = end.UTC().Format("2006-01-02T15:04:05.999Z07:00")
	}
	content, err := r.Request(method, path, params, nil, false)
	if err != nil {
		if _, ok := err.(ErrResponse); ok {
			return nil, err
		}
		return nil, errors.Wrap(err, "candles request")
	}

	var cs []*Candle
	if err := json.Unmarshal(content, &cs); err != nil {
		return nil, errors.Wrap(err, "candles response body")
	}

	return cs, nil
}
