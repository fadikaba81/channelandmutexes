package worker

import (
	"sync"
	"time"

	"github.com/fadikaba81/channelandmutexes/tree/main/internal/result"

)

func Worker(id int, jobs <-chan int, result *result.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		time.sleep(50 * time.Millisecond)
		r.Increment()
	}

}
