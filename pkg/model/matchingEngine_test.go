package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimitBuyOrderHandling(t *testing.T) {
	tests := []struct {
		name                string
		initialSellOrders   []*Order
		buyOrder            Order
		expectedSellOrders  []*Order
		expectedBuyQuantity int
	}{
		{
			name: "Match with sell order",
			initialSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			},
			buyOrder:            Order{ID: 2, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			expectedSellOrders:  []*Order{},
			expectedBuyQuantity: 0,
		},
		{
			name: "Partial match with sell order",
			initialSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 5, Timestamp: time.Now()},
			},
			buyOrder:            Order{ID: 2, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			expectedSellOrders:  []*Order{},
			expectedBuyQuantity: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewMatchingEngine()
			me.OrderBook.SellOrders = tt.initialSellOrders
			me.matchBuyOrder(&tt.buyOrder)

			assertOrdersEqual(t, me.OrderBook.SellOrders, tt.expectedSellOrders)
			assert.Equal(t, tt.expectedBuyQuantity, tt.buyOrder.Quantity)
		})
	}
}

func TestLimitSellOrderHandling(t *testing.T) {
	tests := []struct {
		name                 string
		initialBuyOrders     []*Order
		sellOrder            Order
		expectedBuyOrders    []*Order
		expectedSellQuantity int
	}{
		{
			name: "Match with buy order",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			},
			sellOrder:            Order{ID: 2, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			expectedBuyOrders:    []*Order{},
			expectedSellQuantity: 0,
		},
		{
			name: "Partial match with buy order",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 5, Timestamp: time.Now()},
			},
			sellOrder:            Order{ID: 2, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			expectedBuyOrders:    []*Order{},
			expectedSellQuantity: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewMatchingEngine()
			me.OrderBook.BuyOrders = tt.initialBuyOrders
			me.matchSellOrder(&tt.sellOrder)

			assertOrdersEqual(t, me.OrderBook.BuyOrders, tt.expectedBuyOrders)
			assert.Equal(t, tt.expectedSellQuantity, tt.sellOrder.Quantity)
		})
	}
}

func TestMarketBuyOrderHandling(t *testing.T) {
	tests := []struct {
		name                string
		initialSellOrders   []*Order
		buyOrder            Order
		expectedSellOrders  []*Order
		expectedBuyQuantity int
	}{
		{
			name: "Match with sell order",
			initialSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			},
			buyOrder:            Order{ID: 2, OrderType: Market, Side: Buy, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedSellOrders:  []*Order{},
			expectedBuyQuantity: 0,
		},
		{
			name: "Partial match with sell order",
			initialSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 5, Timestamp: time.Now()},
			},
			buyOrder:            Order{ID: 2, OrderType: Market, Side: Buy, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedSellOrders:  []*Order{},
			expectedBuyQuantity: 5,
		},
		{
			name:                "No match with empty sell orders",
			initialSellOrders:   []*Order{},
			buyOrder:            Order{ID: 2, OrderType: Market, Side: Buy, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedSellOrders:  []*Order{},
			expectedBuyQuantity: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewMatchingEngine()
			me.OrderBook.SellOrders = tt.initialSellOrders
			me.matchBuyOrder(&tt.buyOrder)

			assertOrdersEqual(t, me.OrderBook.SellOrders, tt.expectedSellOrders)
			assert.Equal(t, tt.expectedBuyQuantity, tt.buyOrder.Quantity)
		})
	}
}

func TestMarketSellOrderHandling(t *testing.T) {
	tests := []struct {
		name                 string
		initialBuyOrders     []*Order
		sellOrder            Order
		expectedBuyOrders    []*Order
		expectedSellQuantity int
	}{
		{
			name: "Match with buy order",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: time.Now()},
			},
			sellOrder:            Order{ID: 2, OrderType: Market, Side: Sell, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedBuyOrders:    []*Order{},
			expectedSellQuantity: 0,
		},
		{
			name: "Partial match with buy order",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 5, Timestamp: time.Now()},
			},
			sellOrder:            Order{ID: 2, OrderType: Market, Side: Sell, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedBuyOrders:    []*Order{},
			expectedSellQuantity: 5,
		},
		{
			name:                 "No match with empty buy orders",
			initialBuyOrders:     []*Order{},
			sellOrder:            Order{ID: 2, OrderType: Market, Side: Sell, Price: 0.0, Quantity: 10, Timestamp: time.Now()},
			expectedBuyOrders:    []*Order{},
			expectedSellQuantity: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewMatchingEngine()
			me.OrderBook.BuyOrders = tt.initialBuyOrders
			me.matchSellOrder(&tt.sellOrder)

			assertOrdersEqual(t, me.OrderBook.BuyOrders, tt.expectedBuyOrders)
			assert.Equal(t, tt.expectedSellQuantity, tt.sellOrder.Quantity)
		})
	}
}

func TestGetOrderBook(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name               string
		initialBuyOrders   []*Order
		initialSellOrders  []*Order
		expectedBuyOrders  []*Order
		expectedSellOrders []*Order
	}{
		{
			name:               "Empty order book",
			initialBuyOrders:   []*Order{},
			initialSellOrders:  []*Order{},
			expectedBuyOrders:  []*Order{},
			expectedSellOrders: []*Order{},
		},
		{
			name: "Order book with one buy order",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: now},
			},
			initialSellOrders: []*Order{},
			expectedBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: now},
			},
			expectedSellOrders: []*Order{},
		},
		{
			name:             "Order book with one sell order",
			initialBuyOrders: []*Order{},
			initialSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: now},
			},
			expectedBuyOrders: []*Order{},
			expectedSellOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Sell, Price: 100.0, Quantity: 10, Timestamp: now},
			},
		},
		{
			name: "Order book with multiple orders",
			initialBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: now},
				{ID: 2, OrderType: Limit, Side: Buy, Price: 101.0, Quantity: 5, Timestamp: now},
			},
			initialSellOrders: []*Order{
				{ID: 3, OrderType: Limit, Side: Sell, Price: 102.0, Quantity: 8, Timestamp: now},
				{ID: 4, OrderType: Limit, Side: Sell, Price: 103.0, Quantity: 6, Timestamp: now},
			},
			expectedBuyOrders: []*Order{
				{ID: 1, OrderType: Limit, Side: Buy, Price: 100.0, Quantity: 10, Timestamp: now},
				{ID: 2, OrderType: Limit, Side: Buy, Price: 101.0, Quantity: 5, Timestamp: now},
			},
			expectedSellOrders: []*Order{
				{ID: 3, OrderType: Limit, Side: Sell, Price: 102.0, Quantity: 8, Timestamp: now},
				{ID: 4, OrderType: Limit, Side: Sell, Price: 103.0, Quantity: 6, Timestamp: now},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			me := NewMatchingEngine()
			me.OrderBook.BuyOrders = tt.initialBuyOrders
			me.OrderBook.SellOrders = tt.initialSellOrders

			orderBook := me.GetOrderBook()

			assertOrdersEqual(t, orderBook.BuyOrders, tt.expectedBuyOrders)
			assertOrdersEqual(t, orderBook.SellOrders, tt.expectedSellOrders)
		})
	}
}

func assertOrdersEqual(t *testing.T, actual, expected []*Order) {
	if len(actual) != len(expected) {
		t.Errorf("expected %d orders, got %d orders", len(expected), len(actual))
		return
	}

	for i := range actual {
		assertOrderEqual(t, actual[i], expected[i])
	}
}

func assertOrderEqual(t *testing.T, actual, expected *Order) {
	if actual.ID != expected.ID {
		t.Errorf("expected order ID %d, got %d", expected.ID, actual.ID)
	}
	if actual.OrderType != expected.OrderType {
		t.Errorf("expected order Ordertype %v, got %v", expected.OrderType, actual.OrderType)
	}
	if actual.Side != expected.Side {
		t.Errorf("expected order side %v, got %v", expected.Side, actual.Side)
	}
	if actual.Price != expected.Price {
		t.Errorf("expected order price %v, got %v", expected.Price, actual.Price)
	}
	if actual.Quantity != expected.Quantity {
		t.Errorf("expected order quantity %v, got %v", expected.Quantity, actual.Quantity)
	}
	if !actual.Timestamp.Equal(expected.Timestamp) {
		t.Errorf("expected order timestamp %v, got %v", expected.Timestamp, actual.Timestamp)
	}
}
