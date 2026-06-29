package orderbook

import (
	"github.com/google/btree"
)

// in-memory Orderbook
type Orderbook struct {
	fillIdCounter int64
	askIdCounter  int64
	bidIdCounter  int64
	Asks          *btree.BTreeG[*QuotesWrapper]
	Bids          *btree.BTreeG[*QuotesWrapper]
}

type FillMap map[int64][]Quote // Map<price, []Quote>

func CreateOrderbook() Orderbook {
	orderbookInstance := Orderbook{
		Asks: btree.NewG(3, func(a, b *QuotesWrapper) bool {
			return a.price < b.price
		}),
		Bids: btree.NewG(3, func(a, b *QuotesWrapper) bool {
			return b.price < a.price
		}),
	}

	return orderbookInstance
}

func (orderbook *Orderbook) AddAsk(quote Quote) {
	quoteWrapper, ok := orderbook.Asks.Get(&QuotesWrapper{
		price: quote.Price,
	})
	if !ok {
		orderbook.Asks.ReplaceOrInsert(&QuotesWrapper{
			price:  quote.Price,
			quotes: []Quote{quote},
		})
	} else {
		quoteWrapper.addQuote(quote)
	}
}

func (orderbook *Orderbook) AddBid(quote Quote) {
	quoteWrapper, ok := orderbook.Bids.Get(&QuotesWrapper{
		price: quote.Price,
	})
	if !ok {
		orderbook.Bids.ReplaceOrInsert(&QuotesWrapper{
			price:  quote.Price,
			quotes: []Quote{quote},
		})
	} else {
		quoteWrapper.addQuote(quote)
	}
}

// Returns FillMap of all the asks to fill the quantity@price & quantity of each ask required
func (orderbook *Orderbook) GetAsks(quantity, price int64) FillMap {

	fillMap := make(FillMap)

	orderbook.Asks.Ascend(func(quoteWrapper *QuotesWrapper) bool {

		// If price of asks is more than order price, don't match
		if quoteWrapper.price > price {
			return false
		}

		asks := []Quote{}

		for _, quote := range quoteWrapper.quotes {

			if quantity-quote.Quantity >= 0 {
				asks = append(asks, quote)
				quantity -= quote.Quantity
			} else {
				quote.Quantity = quantity
				asks = append(asks, quote)
				fillMap[quoteWrapper.price] = asks
				return false
			}
		}
		return true
	})

	return fillMap
}

// Returns all the bids to fill the quantity@price & quantity of each bids required
func (orderbook *Orderbook) GetBids(quantity, price int64) FillMap {

	fillMap := make(FillMap)

	orderbook.Bids.Ascend(func(quoteWrapper *QuotesWrapper) bool {

		// If price of bids is less than order price, don't match
		if quoteWrapper.price < price {
			return false
		}

		bids := []Quote{}

		for _, quote := range quoteWrapper.quotes {

			if quantity-quote.Quantity >= 0 {
				bids = append(bids, quote)
				quantity -= quote.Quantity
			} else {
				quote.Quantity = quantity
				bids = append(bids, quote)
				fillMap[quoteWrapper.price] = bids
				return false
			}
		}
		return true
	})

	return fillMap
}

func (orderbook *Orderbook) DeleteAsks(fillMap FillMap) {
	orderbook.Asks.Ascend(func(quoteWrapper *QuotesWrapper) bool {

		if len(fillMap[quoteWrapper.price]) <= 0 {
			return false
		}

		for _, quote := range fillMap[quoteWrapper.price] {
			quoteWrapper.removeQuote(quote)
		}
		return true
	})
}

func (orderbook *Orderbook) DeleteBids(fillMap FillMap) {
	orderbook.Bids.Ascend(func(quoteWrapper *QuotesWrapper) bool {

		if len(fillMap[quoteWrapper.price]) <= 0 {
			return false
		}

		for _, quote := range fillMap[quoteWrapper.price] {
			quoteWrapper.removeQuote(quote)
		}
		return true
	})
}

func (orderbook *Orderbook) GetFillCounter() int64 {
	orderbook.fillIdCounter++
	return orderbook.fillIdCounter
}

func (orderbook *Orderbook) GetAskCounter() int64 {
	orderbook.askIdCounter++
	return orderbook.askIdCounter
}

func (orderbook *Orderbook) GetBidCounter() int64 {
	orderbook.bidIdCounter++
	return orderbook.bidIdCounter
}
