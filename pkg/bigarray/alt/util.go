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
	defer os.Remove(f.Name())

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

func haveNeed(h, n int) string {
	return fmt.Sprintf("Have: %d; need: %d", h, n)
}
