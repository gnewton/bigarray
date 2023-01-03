package alt

import (
	"fmt"
	"log"
	"os"
	"time"
)

func tmpFile(tmpDir string) (*os.File, error) {
	if len(tmpDir) == 0 {
		tmpDir = "."
	}

	now := time.Now()
	tim := fmt.Sprintf("%d%s%s-%s%s%s", now.Year(), zero(int(now.Month())), zero(now.Day()), zero(now.Hour()), zero(now.Minute()), zero(now.Second()))

	f, err := os.CreateTemp(".", SeqFilePrefix+"_"+tim+"__")
	if err != nil {
		return nil, err
	}
	log.Println("Opened: ", f.Name())
	return f, nil
}

// Prefixes zeros to 0-9 ints
func zero(i int) string {
	v := fmt.Sprintf("%d", i)
	if i > 9 {
		return v
	}
	return "0" + v
}

func haveNeed(h, n int64) string {
	return fmt.Sprintf("Have: %d; need: %d", h, n)
}

func notImplemented() error {
	return fmt.Errorf("Not implemented")
}

func shouldBeError() error {
	return fmt.Errorf("Not implemented")
}
