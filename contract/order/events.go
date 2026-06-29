package order

type TradeType int
type TradeAction int

const (
	Market TradeType = iota + 1
	Limit
)

const (
	Buy TradeAction = iota + 1
	Sell
)

type CreateOrder struct {
	Id          string
	UserId      int32
	MarketId    int32
	TradeType   TradeType
	TradeAction TradeAction
	Price       int64
	Quantity    int64
	IsIOC       bool
	IsPostOnly  bool
}

type CancelOrder struct {
	Id string
}
