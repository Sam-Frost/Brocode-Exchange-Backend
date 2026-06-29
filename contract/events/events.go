package event

type Event[T any] struct {
	Type string
	Data T
}

const (
	CreateSpotOrder string = "create-spot-order"
	CreatePerpOrder string = "create-perp-order"
	CancelOrder     string = "cancel-order"
)
