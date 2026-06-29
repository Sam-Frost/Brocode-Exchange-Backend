package orderbook

import "slices"

// Represents all the quotes(ask/bid) at the price
type QuotesWrapper struct {
	price  int64
	quotes []Quote
}

func (q *QuotesWrapper) addQuote(quote Quote) {
	q.quotes = append(q.quotes, quote)
}

func (q *QuotesWrapper) removeQuote(quote Quote) {
	for index, currQuote := range q.quotes {
		if currQuote.Id == quote.Id {
			if currQuote.Quantity == quote.Quantity {
				q.quotes = slices.Delete(q.quotes, index, index+1)
			} else {
				q.quotes[index].Quantity -= quote.Quantity
			}
			return
		}
	}
}

// Represents a single quote(ask/bid)
type Quote struct {
	Id       int64
	OrderId  string
	UserId   int32
	Quantity int64
	Price    int64
}
