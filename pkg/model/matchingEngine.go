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

/*-------------------------------------------
 MATCH ENGINE METHODS
-------------------------------------------*/
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

	// logic to set up the order book within the match engine
	if orderType == Market {
		me.matchOrder(order)
	} else {

		if side == Buy {
			// add buy order to the buy side order book
			me.OrderBook.BuyOrders = append(me.OrderBook.BuyOrders, order)
		} else {
			// add sell order to the sell side order book
			me.OrderBook.SellOrders = append(me.OrderBook.SellOrders, order)
		}
		me.matchOrder(order)
	}

	return order
}

// matchOrder matches an order against the order book
func (me *MatchingEngine) matchOrder(order *Order) {
	if order.Side == Buy {
		// Match buy order
		me.matchBuyOrder(order)
	} else {
		// Match sell order
		me.matchSellOrder(order)
	}
}

// function matches buy side orders
func (me *MatchingEngine) matchBuyOrder(order *Order) {
	// if order book is not empty and the order has been made at least once
	for len(me.OrderBook.SellOrders) > 0 && order.Quantity > 0 {

		bestSellOrder := me.OrderBook.SellOrders[0]
		if order.OrderType == Limit && order.Price < bestSellOrder.Price {
			break
		}
		if order.Quantity >= bestSellOrder.Quantity {
			order.Quantity -= bestSellOrder.Quantity
			me.OrderBook.SellOrders = me.OrderBook.SellOrders[1:]
		} else {
			bestSellOrder.Quantity -= order.Quantity
			order.Quantity = 0
		}
	}
}

// function matches sell side orders
func (me *MatchingEngine) matchSellOrder(order *Order) {
	// if order book is not empty and the order has been made at least once
	for len(me.OrderBook.BuyOrders) > 0 && order.Quantity > 0 {
		bestBuyOrder := me.OrderBook.BuyOrders[0]
		if order.OrderType == Limit && order.Price > bestBuyOrder.Price {
			break
		}
		if order.Quantity >= bestBuyOrder.Quantity {
			order.Quantity -= bestBuyOrder.Quantity
			me.OrderBook.BuyOrders = me.OrderBook.BuyOrders[1:]
		} else {
			bestBuyOrder.Quantity -= order.Quantity
			order.Quantity = 0
		}
	}
}

// GetOrderBook returns the current state of the order book
func (me *MatchingEngine) GetOrderBook() *OrderBook {
	return me.OrderBook
}
