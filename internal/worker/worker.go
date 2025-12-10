package worker

import (

	"sync"
	"time"

	"github.com/fadikaba81/channelandmutexes/internal/result"
)

func Worker(id int, jobs <-chan int, r *result.Result, wg *sync.WaitGroup) {

	defer wg.Done()


	for _ = range jobs {
		time.Sleep(50 * time.Millisecond)
		r.Increment()
	}

}
