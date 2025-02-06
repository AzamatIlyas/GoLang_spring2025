package main

import (
	"fmt"
	"time"
)


func main(){
	t := time.Now()
	const jobsCount, workerCount = 15, 5
	jobs := make(chan int, 15)
	res := make(chan int, 15)

	for i := 0; i < workerCount; i++{
		go worker(i+1, jobs, res)
	}

	for i := 0; i < jobsCount; i++ {
		jobs <- i+1
	}
	close(jobs)

	for i := 0; i < jobsCount; i++{
		fmt.Printf("result #%d : value = %d\n", i+1, <-res)
	}
	fmt.Println(time.Since(t))
}



func worker(id int, jobs <-chan int, res chan<- int ) {
	for j := range jobs {
		time.Sleep(time.Second)
		fmt.Printf("worker #%d finished\n", id)
		res <- j*j
	}
}