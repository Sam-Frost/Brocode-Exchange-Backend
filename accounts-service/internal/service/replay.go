package service

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"

	"github.com/Sam-Frost/accounts-service/internal/db"
)

func getSnapshotDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Reading home dir: %s", err)
	}

	return fmt.Sprintf("%s/.accounts-service/snapshot.bin", homeDir)
}

func RecreateDatabase() {
	file, _ := os.Open(getSnapshotDir())

	decoder := gob.NewDecoder(file)

	decoder.Decode(db.GetDatabase())
	file.Close()
}
