package model

import "time"

// MatchingEngine represents the matching engine
type MatchingEngine struct {
	OrderBook *OrderBook
	OrderID   int
}

// NewMatchingEngine creates a new matching engine
func NewMatchingEngine() *MatchingEngine {
	return &MatchingEngine{
		OrderBook: &OrderBook{
			BuyOrders:  []*Order{},
			SellOrders: []*Order{},
		},
		OrderID: 0,
	}
}

// PlaceOrder places a new order in the matching engine
func (me *MatchingEngine) PlaceOrder(orderType OrderType, side OrderSide, price float64, quantity int) *Order {
	// increment order id for a new order
	me.OrderID++
	// create a new order object
	order := &Order{
		ID:        me.OrderID,
		OrderType: orderType,
		Side:      side,
		Price:     price,
		Quantity:  quantity,
		Timestamp: time.Now(),
	}

	return order
}
