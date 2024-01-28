package persistance

import (
	"io/ioutil"
	"os"

	"testing"

	"github.com/hitesharora1997/testassignment/internal/counter"
)

func TestPersistence_SaveAndRestore(t *testing.T) {
	// Create a temporary file
	tmpFile, err := ioutil.TempFile("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	c := counter.NewRequestCounter()

	// Edge Case 1: Save and restore with no data
	SaveData(tmpFile.Name(), c)
	newCounter := counter.NewRequestCounter()
	if err = Restore(tmpFile.Name(), newCounter); err != nil {
		t.Errorf("Restore failed with no data: %v", err)
	}

	// Edge Case 2: Save and restore with some data
	c.RecordAndCount()
	SaveData(tmpFile.Name(), c)
	newCounter = counter.NewRequestCounter()
	Restore(tmpFile.Name(), newCounter)
	if len(newCounter.RequestTimes) != 1 {
		t.Errorf("Expected 1 request, got %d", len(newCounter.RequestTimes))
	}

	// Edge Case 3: Handling file read/write errors (e.g., permission issues, corrupted data)
	// This can be simulated by using an invalid file path or modifying file permissions.
}
