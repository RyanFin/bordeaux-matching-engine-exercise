package model

import "time"

type OrderType string

const (
	Market OrderType = "market"
	Limit  OrderType = "limit"
)

// OrderSide represents the side of an order (buy or sell)
type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

// Trading Order
type Order struct {
	ID        int       `json:"id"`
	OrderType OrderType `json:"type"`
	Side      OrderSide `json:"side"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
}

// OrderBook represents the order book with buy and sell orders
type OrderBook struct {
	BuyOrders  []*Order
	SellOrders []*Order
}
