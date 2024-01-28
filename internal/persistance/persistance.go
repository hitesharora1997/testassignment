package persistance

import (
	"encoding/json"
	"os"

	"github.com/hitesharora1997/testassignment/internal/counter"
)

func Restore(dataFile string, counter *counter.RequestCounter) error {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return err
	}

	counter.Lock()
	defer counter.Unlock()

	return json.Unmarshal(data, counter.RecordCounter())
}
