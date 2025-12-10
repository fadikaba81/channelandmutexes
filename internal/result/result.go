package result

import "sync"

type Result struct {
	mu   sync.Mutex
	count int
}


func New() *Result {
return &Result{}
}

func (r *Result) Increment() {

	r.mu.Lock()
	r.count++
	r.mu.Unlock()
}

func (r *Result) GetValue() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.count
}