package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("第%d个goroutine开始处理编号为%d的job\n", id, job)
		results <- job * 2
		time.Sleep(500 * time.Microsecond)
		fmt.Printf("第%d个goroutine处理结束编号为%d的job\n", id, job)

	}

}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//使用3个goroutine来处理任务
	for i := 0; i < 30; i++ {
		go worker(i, jobs, results)
	}

	//发送5个任务
	for i := 0; i < 50; i++ {
		jobs <- i
	}
	close(jobs)

	//打印results
	for i := 0; i < 50; i++ {
		fmt.Println(<-results)
	}

}
