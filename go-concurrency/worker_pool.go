package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Println("worker", id, "processing job", job)
		time.Sleep(time.Second)
		results <- job * 2
	}

}
func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for i := 1; i <= 5; i++ {
		fmt.Println("Results:", <-results)
	}
}
