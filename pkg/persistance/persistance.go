package persistance

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/hitesharora1997/testassignment/internal/counter"
)

func Restore(dataFile string, counter *counter.RequestCounter) error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	counter.Lock()
	defer counter.Unlock()

	return json.Unmarshal(data, counter.RecordAndCount())
}

func PeriodicSave(dataFile string, counter *counter.RequestCounter, intervalSeconds int) {
	for {
		time.Sleep(time.Duration(intervalSeconds) * time.Second)
		SaveData(dataFile, counter)
	}
}

func SaveData(dataFile string, counter *counter.RequestCounter) {
	counter.Lock()
	defer counter.Unlock()

	data, err := json.Marshal(counter.RequestTimes)
	if err != nil {
		log.Println("Error marshalling the data:", err)
	}

	// read and write the file and other's can only read it
	if err = os.WriteFile(dataFile, data, 0644); err != nil {
		log.Println("Error writing file:", err)
	}
}
