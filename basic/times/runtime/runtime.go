package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano()
	//fmt.Printf("start time is:", start.Format("2016-01-02 15:04:00"))
	for i := 0; i < 2; i++ {
		time.Sleep(1 * time.Second)
	}
	end := time.Now().UnixNano()
	//fmt.Printf("end time is:", end.Format("2016-01-02 15:04:00"))
	t := end - start
	fmt.Printf("usetime %d nanosecond", t)
}
