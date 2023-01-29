package solver

import (
	"log"
	"time"
)

// Runningtime computes running time
func Runningtime(s string) (string, time.Time) {
	log.Println("Start: ", s)
	return s, time.Now()
}

// Track is this
func Track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:   ", s, "took", endTime.Sub(startTime))
}
