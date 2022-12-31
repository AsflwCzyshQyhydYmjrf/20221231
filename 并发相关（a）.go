package awesomeProject1

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	go func() {
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()
	mu.Unlock()
}
