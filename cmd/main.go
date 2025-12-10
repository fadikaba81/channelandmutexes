package main

import (
	"fmt"
	"sync"

	"github.com/fadikaba81/channelandmutexes/tree/main/internal/result"
	"github.com/fadikaba81/channelandmutexes/tree/main/internal/worker"
)

func main() {

	job := make(chan int)
	var wg sync.WaitGroup

	result := result.New()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker.Worker(i, job, result, &wg)
	}

	for i := 0; i < 10; i++ {
		job <- i
	}
	close(job)
	wg.Wait()

	fmt.Println("Final Results:", result.GetAll())
}
