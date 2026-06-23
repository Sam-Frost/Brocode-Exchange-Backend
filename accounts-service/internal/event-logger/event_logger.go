package eventLogger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const EVENT_CHANNEL_BUFFER int = 100000
const (
	APP_HOME_DIR        = ".accounts-service"
	LOG_FILE            = "event.log"
	OFFSET_TRACKER_FILE = "offset_tracker.json"
)

type OrderData struct {
}

type Event struct {
	EventId  EventType
	UserId   int32
	MarketId int32
	Amount   int64
	Sequence int64
	OrderData
	timestamp string
}

type offsetData struct {
	LineNumber int64
}

// Channel to produce event too
var eventChannel = make(chan Event, EVENT_CHANNEL_BUFFER)

func InitEventLogFile() {

	filePath := getLogFilePath()
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		if !os.IsExist(err) {
			log.Fatalf("Error creating log file: %v", err)
		}
	}
	file.Close()

}

// Attach UTC timestamp and emit event to channel
func EmitEvent(eventData Event) {
	eventData.timestamp = time.Now().In(time.UTC).String()
	eventChannel <- eventData
}

func InitEventLogger() {
	filePath := getLogFilePath()
	eventFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {

	}
	defer eventFile.Close()
	for eventData := range eventChannel {

		jsonData, err := json.Marshal(eventData)
		if err != nil {

		}
		jsonData = append(jsonData, byte('\n'))
		eventFile.Write(jsonData)
	}
}

func InitOffsetTracker() {
	filePath := getOffsetTackerFilePath()
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return
		} else {
			log.Fatalf("Error creating offset tracker file: %v", err)
		}
	}
	defer file.Close()

	initOffest := offsetData{
		LineNumber: 1,
	}

	fmt.Print(initOffest)
	bytesSlice, err := json.Marshal(initOffest)
	if err != nil {
		log.Fatalf("Error converting init offset data to bytes: %v", err)
	}

	_, err = file.Write(bytesSlice)
	if err != nil {
		log.Fatalf("Error occured while writing to offset_tracker: %v", err)
	}
}

func SendEventToKafkaBroker() {
	offset := getCurrentOffset()
	fmt.Println(offset)

	// logFile, _ := os.Re
	// Read Events from log file
	// Send to Kafka
	// Write to file the row sent
	//

}

func getOffsetTackerFilePath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to get user home dir: %v", err)
	}

	return fmt.Sprintf("%s/%s/%s", userHomeDir, APP_HOME_DIR, OFFSET_TRACKER_FILE)
}

func getLogFilePath() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Unable to get user home dir: %v", err)
	}

	return fmt.Sprintf("%s/%s/%s", userHomeDir, APP_HOME_DIR, LOG_FILE)
}

func getCurrentOffset() int64 {
	filePath := getOffsetTackerFilePath()
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file : %v", err)
	}

	var offset offsetData

	err = json.Unmarshal(fileBytes, &offset)
	if err != nil {
		log.Fatalf("Error parsing offset_tracker data: %v", err)
	}

	return offset.LineNumber
}
