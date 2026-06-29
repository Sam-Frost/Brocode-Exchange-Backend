package services

import (
	"fmt"

	"github.com/Sam-Frost/contract/order"
	"github.com/Sam-Frost/matching-engine/internal/orderbook"
	ob "github.com/Sam-Frost/matching-engine/internal/orderbook"
	"github.com/Sam-Frost/matching-engine/internal/util"
)

type Fill struct {
	id           int64
	takerOrderId string
	takerUserId  int32
	makerOrderId string
	makerUserId  int32
	quantity     int64
	price        int64
}

func CreateOrder(orderData order.CreateOrder, orderbook *ob.Orderbook, fillsRingBuf *util.RingBuffer[[]Fill]) {
	switch orderData.TradeAction {
	case order.Buy:
		{
			fillMap := orderbook.GetAsks(orderData.Quantity, orderData.Price)
			fills := generateFills(orderData, fillMap, orderbook)
			orderbook.DeleteAsks(fillMap)

			if orderData.TradeType == order.Limit && !orderData.IsIOC {

				var totalFilledQuanity int64
				for _, quotes := range fillMap {
					for _, quote := range quotes {
						totalFilledQuanity += quote.Quantity
					}
				}

				if totalFilledQuanity != orderData.Quantity {
					orderbook.AddBid(ob.Quote{
						Id:       orderbook.GetBidCounter(),
						OrderId:  orderData.Id,
						UserId:   orderData.UserId,
						Quantity: orderData.Quantity - totalFilledQuanity,
						Price:    orderData.Price,
					})
				}
			}

			fillsRingBuf.PushWait(fills)
		}

	case order.Sell:
		{

		}
	default:
		fmt.Println("Invalid trade action")
		return

	}

	// Send fills to lock free ring

}

func CancelOrder() {

}

func generateFills(orderData order.CreateOrder, fillMap orderbook.FillMap, orderbook *orderbook.Orderbook) []Fill {
	fills := []Fill{}

	for _, quotes := range fillMap {
		for _, quote := range quotes {
			fill := Fill{
				id:           orderbook.GetFillCounter(),
				takerOrderId: orderData.Id,
				takerUserId:  orderData.UserId,
				makerOrderId: quote.OrderId,
				makerUserId:  quote.UserId,
				quantity:     quote.Quantity,
				price:        quote.Price,
			}
			fills = append(fills, fill)
		}
	}

	return fills
}

func SendFillsToProducer(fills []Fill) {

}
