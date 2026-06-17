package eventLogger

import "fmt"

type Event struct {
	EventId  EventType
	UserId   int32
	Amount   int64
	MarketId int32
	Sequence int64
}

// Channel to produce event too
var eventChannel = make(chan Event)

func EmitEvent(eventData Event) {
	eventChannel <- eventData
}

func InitEventLogger() {
	for eventData := range eventChannel {
		// TODO : Persiste data into a file
		// TODO : Snapshot the accounts book every 5 mins
		fmt.Println(eventData)
	}
}
