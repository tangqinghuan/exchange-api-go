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

// Account ...
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

// Bill ...
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

// Candle ...
type Candle struct {
	// Start time
	Time time.Time `json:"time"`
	// 	Lowest price
	Low decimal.Decimal `json:"low"`
	// Highest price
	High decimal.Decimal `json:"high"`
	// Open price
	Open decimal.Decimal `json:"open"`
	// Close price
	Close decimal.Decimal `json:"close"`
	// 	Trading volume
	Volume decimal.Decimal `json:"volume"`
}
