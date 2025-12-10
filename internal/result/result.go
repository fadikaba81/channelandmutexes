package result

import "sync"

type Result struct {
	mu   sync.Mutex
	data map[int]int
}
