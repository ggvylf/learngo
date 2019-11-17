package main

import (
	"fmt"
	"time"
)

func mytimer() {
	timer1 := time.NewTimer(3 * time.Second)

	fmt.Println("timer1 start=", time.Now())
	<-timer1.C
	fmt.Println("timer1 end=", time.Now())

	timer2 := time.After(3 * time.Second)
	fmt.Println("timer2 start=", time.Now())
	<-timer2
	fmt.Println("timer2 end=", time.Now())

	fmt.Println("timer3 start=", time.Now())
	<-time.After(3 * time.Second)
	fmt.Println("timer3 end=", time.Now())

}

func stoptimer() {
	fmt.Println("start=", time.Now())
	timer := time.NewTimer(5 * time.Second)
	time.Sleep(3 * time.Second)
	//终止计时
	ok := timer.Stop()
	if ok {
		fmt.Println("stop ok")
	}
	//stop以后，timer.C里没有数据，读取会panic
	// end := <-timer.C
	// fmt.Println("end=", end)
	fmt.Println("end=", time.Now())
}

func resettimer() {
	fmt.Println("start=", time.Now())
	timer := time.NewTimer(2 * time.Second)

	//reset要放在timer.C写入之前
	ok := timer.Reset(2 * time.Second)
	if ok {
		fmt.Println("reset ok")
	}

	time.Sleep(3 * time.Second)
	end := <-timer.C
	fmt.Println("end=", end)
}

func main() {

	// mytimer()
	stoptimer()
	// resettimer()
}
