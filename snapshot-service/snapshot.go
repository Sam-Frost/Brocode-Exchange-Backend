package service

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sam-Frost/accounts-service/internal/db"
)

const DATABASE_BACKUP_INTERVAL int = 15 // Mins

func BackupDatabase() {
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		SnapshotDatabase()
	}
}

func getSnapshotDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Reading home dir: %s", err)
	}

	return fmt.Sprintf("%s/.accounts-service/snapshot.bin", homeDir)
}

func SnapshotDatabase() {
	file, err := os.Create(getSnapshotDir())
	if err != nil {
		fmt.Println(err)
	}
	encoder := gob.NewEncoder(file)

	encoder.Encode(db.GetDatabase())

}

// func PrettyPrintDatabase() {
// 	db := db.GetDatabase()

// 	for key := range *db {
// 		fmt.Printf("Key : %v | ", key)

// 		userData := (*(*db)[key])
// 		balanceData := *(userData.BalanceData)
// 		fmt.Printf("Balance Data : %v \n", balanceData)
// 	}
// }
