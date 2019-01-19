package okex_test

import (
	"testing"
	"time"

	"github.com/wanluc/exchange-api-go/okex/v3"
)

var rest = okex.NewRest("")

func TestCandles(t *testing.T) {
	end := time.Now()
	start := end.Add(-5 * time.Minute)
	cs, err := rest.Candles("BTC-USDT", 60, &start, &end)
	if err != nil {
		t.Fatal(err)
	}

	for _, c := range cs {
		t.Logf("candle : %v", c)
	}
}

func TestInstruments(t *testing.T) {
	ss, err := rest.Instruments()
	if err != nil {
		t.Fatal(err)
	}

	for _, s := range ss {
		t.Logf("instrument : %v", s)
	}
}

func TestOrderBook(t *testing.T) {
	ob, err := rest.OrderBook("BTC-USDT", "", 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, b := range ob.Bids {
		t.Logf("order book bid : %v", b)
	}
	for _, a := range ob.Asks {
		t.Logf("order book ask : %v", a)
	}
}

func TestAllTicker(t *testing.T) {
	ts, err := rest.AllTicker()
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range ts {
		t.Logf("ticker : %v", r)
	}
}

func TestTicker(t *testing.T) {
	tk, err := rest.Ticker("BTC-USDT")
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("ticker : %v", tk)
}

func TestTrades(t *testing.T) {
	ts, err := rest.Trades("BTC-USDT", "", "", 10)
	if err != nil {
		t.Fatal(err)
	}

	for _, r := range ts {
		t.Logf("trade : %v", r)
	}
}
