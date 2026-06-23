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
		if currQuote.id == quote.id {
			if currQuote.quantity == quote.quantity {
				q.quotes = slices.Delete(q.quotes, index, index+1)
			} else {
				q.quotes[index].quantity -= quote.quantity
			}
			return
		}
	}
}

// Represents a single quote(ask/bid)
type Quote struct {
	id               int64
	orderId          int32
	userId           int32
	quantity         int64
	price            int64
	lockedCollateral int64
}
