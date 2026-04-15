package availability

import (
	"fmt"
	"time"
)

func NextQuarter(t time.Time) string {
	month := int(t.Month())
	year := t.Year()

	currentQuarter := (month-1)/3 + 1

	nextQuarter := currentQuarter + 1
	nextYear := year

	if nextQuarter > 4 {
		nextQuarter = 1
		nextYear++
	}

	return fmt.Sprintf("Q%d %d", nextQuarter, nextYear)
}
