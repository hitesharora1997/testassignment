package persistance

import (
	"io/ioutil"
	"os"

	"testing"

	"github.com/hitesharora1997/testassignment/internal/counter"
)

func TestPersistence_SaveAndRestore(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	c := counter.NewRequestCounter()

	SaveData(tmpFile.Name(), c)
	newCounter := counter.NewRequestCounter()
	if err = Restore(tmpFile.Name(), newCounter); err != nil {
		t.Errorf("Restore failed with no data: %v", err)
	}

	c.RecordAndCount()
	SaveData(tmpFile.Name(), c)
	newCounter = counter.NewRequestCounter()
	Restore(tmpFile.Name(), newCounter)
	if len(newCounter.RequestTimes) != 1 {
		t.Errorf("Expected 1 request, got %d", len(newCounter.RequestTimes))
	}

}
