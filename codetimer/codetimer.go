package codetimer

import (
	"fmt"
	"time"
)

// CodeTime : For saving time data
type CodeTime struct {
	StartTime time.Time
	EndTime   time.Time
}

// Start : Set StartTime
func (ct *CodeTime) Start() {
	ct.StartTime = time.Now()
}

// End : Set EndTime
func (ct *CodeTime) End() {
	ct.EndTime = time.Now()
	fmt.Println("Code Running time: ", ct.EndTime.Sub(ct.StartTime))
}
