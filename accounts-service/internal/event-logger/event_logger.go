package eventLogger

import (
	"encoding/json"
	"os"
	"time"
)

const EVENT_CHANNEL_BUFFER int = 100

type Event struct {
	EventId   EventType
	UserId    int32
	MarketId  int32
	Amount    int64
	Sequence  int64
	timestamp string
}

// Channel to produce event too
var eventChannel = make(chan Event, EVENT_CHANNEL_BUFFER)

// Attach UTC timestamp and emit event to channel
func EmitEvent(eventData Event) {
	eventData.timestamp = time.Now().In(time.UTC).String()
	eventChannel <- eventData
}

func InitEventLogger() {
	for eventData := range eventChannel {

		eventFile, err := os.OpenFile("event.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {

		}
		defer eventFile.Close()

		jsonData, err := json.Marshal(eventData)
		jsonData = append(jsonData, byte('\n'))
		eventFile.Write(jsonData)
	}
}
