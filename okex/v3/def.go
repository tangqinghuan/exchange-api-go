package okex

import (
	"time"

	"github.com/shopspring/decimal"
)

const (
	// OrderTypeLimit ...
	OrderTypeLimit = "limit"
	// OrderTypeMarket ...
	OrderTypeMarket = "market"

	// OrderSideBuy ...
	OrderSideBuy = "buy"
	// OrderSideSell ...
	OrderSideSell = "sell"
)

// Account represents the balance, amount available/on hold of a token in spot account.
type Account struct {
	// token
	Currency string `json:"currency"`
	// balance
	Balance decimal.Decimal `json:"balance"`
	// amount on hold(not available)
	Hold decimal.Decimal `json:"hold"`
	// available amount
	Available decimal.Decimal `json:"available"`
	// account ID
	ID string `json:"id"`
}

// Bill represents ledger of a token in spot account.
type Bill struct {
	// bill ID
	LedgerID string `json:"ledger_id"`
	// balance
	Balance decimal.Decimal `json:"balance"`
	// token
	Currency string `json:"currency"`
	// amount
	Amount decimal.Decimal `json:"amount"`
	// type of bills
	Type string `json:"type"`
	// creation time
	Timestamp time.Time `json:"timestamp"`
	// creation time
	CreatedAt time.Time `json:"created_at"`
	// details	string	if the type is trade or fee, order details will be included under this
	Details struct {
		OrderID      string `json:"order_id"`
		InstrumentID string `json:"instrument_id"`
		ProductID    string `json:"product_id"`
	} `json:"details"`
}

// OrderRequest represents new order request data.
type OrderRequest struct {
	// client_oid [optional]the order ID customized by yourself
	ClientOID string `json:"client_oid"`
	// type [required]limit / market(default: limit)
	Type string `json:"type"`
	// side [required]buy or sell
	Side string `json:"side"`
	// instrument_id [required]trading pair
	InstrumentID string `json:"instrument_id"`
	// margin_trading [required]order type (1 spot order. 2 margin order)
	MarginTrading int8 `json:"margin_trading"`
	// price [required]	price
	Price string `json:"price"`
	// limit order : size [required]quantity bought or sold
	// market order : size [required]quantity sold. (for orders sold at market price only)
	Size string `json:"size"`
	// market order : notional [required]amount bought. (for orders bought at market price only)
	Notional string `json:"notional"`
}

// OrderResponse represents new order response data.
type OrderResponse struct {
	Result    bool   `json:"result"`
	OrderID   string `json:"order_id"`
	ClientOid string `json:"client_oid"`
}

// BatchCancelOrderRequest represents batch cancel order request data.
type BatchCancelOrderRequest struct {
	// by providing this parameter, the corresponding order of a designated trading pair will be cancelled. If not providing this parameter, it will be back to error code.
	InstrumentID string `json:"instrument_id"`
	// order ID. You may cancel up to 4 orders of a trading pair
	OrderIDs []int64 `json:"order_ids"`
}

// BatchCancelOrderResponse represents batch cancel order response data.
type BatchCancelOrderResponse struct {
	Result    bool     `json:"result"`
	OrderID   []string `json:"order_id"`
	ClientOid string   `json:"client_oid"`
}

// Order represents order data.
type Order struct {
	// order ID
	OrderID string `json:"order_id"`
	// price
	Price decimal.Decimal `json:"price"`
	// quantity
	Size decimal.Decimal `json:"size"`
	// the total buying amount. This value will be returned for market orders
	Notional string `json:"notional"`
	// trading pair
	InstrumentID string `json:"instrument_id"`
	// limit,market(defaulted as limit)
	Type string `json:"type"`
	// buy or sell
	Side string `json:"side"`
	// create date
	Timestamp time.Time `json:"timestamp"`
	// creation time
	CreatedAt time.Time `json:"created_at"`
	// quantity filled
	FilledSize decimal.Decimal `json:"filled_size"`
	// amount filled
	FilledNotional decimal.Decimal `json:"filled_notional"`
	// order status
	// open = new, process = processing, done = completed, cancel = cancelled
	Status        string          `json:"status"`
	ExecutedValue decimal.Decimal `json:"executed_value"`
	Funds         string          `json:"funds"`
	ProductID     string          `json:"product_id"`
}

// Fill represents detail of filled order.
type Fill struct {
	// bill ID
	LedgerID string `json:"ledger_id"`
	// trading pair
	InstrumentID string `json:"instrument_id"`
	// price
	Price decimal.Decimal `json:"price"`
	// quantity
	Size decimal.Decimal `json:"size"`
	// order ID
	OrderID string `json:"order_id"`
	// create date
	Timestamp time.Time `json:"timestamp"`
	// creation time
	CreatedAt time.Time `json:"created_at"`
	// liquidity side (T or M)
	ExecType string `json:"exec_type"`
	// fee amount
	Fee decimal.Decimal `json:"fee"`
	// bills side(buy ,sell or points_fee)
	Side string `json:"side"`
	// liquidity side (T or M)
	Liquidity string `json:"liquidity"`
	ProductID string `json:"product_id"`
}

// Instrument represents token pair.
type Instrument struct {
	// trading pair
	InstrumentID string `json:"instrument_id"`
	// base currency
	BaseCurrency string `json:"base_currency"`
	// quote currency
	QuoteCurrency string `json:"quote_currency"`
	// minimum trading size
	MinSize decimal.Decimal `json:"min_size"`
	// minimum increment size
	SizeIncrement decimal.Decimal `json:"size_increment"`
	// trading price increment
	TickSize decimal.Decimal `json:"tick_size"`

	ProductID      string          `json:"product_id"`
	BaseIncrement  decimal.Decimal `json:"base_increment"`
	BaseMinSize    decimal.Decimal `json:"base_min_size"`
	QuoteIncrement decimal.Decimal `json:"quote_increment"`
}

// OrderBook represents order book of a trading pair.
type OrderBook struct {
	// create date
	Timestamp time.Time `json:"timestamp"`
	// [ price, size, num_orders ]
	Bids [][]decimal.Decimal `json:"bids"`
	Asks [][]decimal.Decimal `json:"asks"`
}

// Ticker represents info of all trading pairs.
type Ticker struct {
	// trading pair
	InstrumentID string `json:"instrument_id"`
	// last traded price
	Last decimal.Decimal `json:"last"`
	// best bid price
	BestBid decimal.Decimal `json:"best_bid"`
	// best ask price
	BestAsk decimal.Decimal `json:"best_ask"`
	// 24 hour open
	Open24H decimal.Decimal `json:"open_24h"`
	// 24 hour high
	High24H decimal.Decimal `json:"high_24h"`
	// 24 hour low
	Low24H decimal.Decimal `json:"low_24h"`
	// 24 trading volume of the base currency
	BaseVolume24H decimal.Decimal `json:"base_volume_24h"`
	// 24 trading volume of the quote currency
	QuoteVolume24H decimal.Decimal `json:"quote_volume_24h"`
	// creation time
	Timestamp time.Time `json:"timestamp"`

	ProductID string `json:"product_id"`
	// bid price
	Bid decimal.Decimal `json:"bid"`
	// ask price
	Ask decimal.Decimal `json:"ask"`
}

// Trade represents transactions of a trading pair.
type Trade struct {
	// filled time
	Timestamp time.Time `json:"timestamp"`
	Time      time.Time `json:"time"`
	// transaction time ID
	TradeID string `json:"trade_id"`
	// filled price
	Price decimal.Decimal `json:"price"`
	// filled size
	Size decimal.Decimal `json:"size"`
	// filled side
	Side string `json:"side"`
}

// Candle represents charts of a trading pair.
type Candle struct {
	// start time
	Time time.Time `json:"time"`
	// lowest price
	Low decimal.Decimal `json:"low"`
	// highest price
	High decimal.Decimal `json:"high"`
	// open price
	Open decimal.Decimal `json:"open"`
	// close price
	Close decimal.Decimal `json:"close"`
	// trading volume
	Volume decimal.Decimal `json:"volume"`
}
