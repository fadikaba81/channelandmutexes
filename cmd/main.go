package main

import (
	"fmt"
	"sync"

	"github.com/fadikaba81/channelandmutexes/internal/result"
	"github.com/fadikaba81/channelandmutexes/internal/worker"
)

func main() {

	job := make(chan int)
	var wg sync.WaitGroup
	var i int

	result := result.New()

	for i = 0; i < 3; i++ {
		wg.Add(1)
		go worker.Worker(i, job, result, &wg)
	}

	for i = range 10 {
		fmt.Println("Sending job", i)
		job <- i
	}
	close(job)
	wg.Wait()

	fmt.Println("Final Results:", result.GetValue())
}
