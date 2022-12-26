package mutex

import (
	"fmt"
	"sync"
)

var (
	balance int = 0
)

// Writer (just can be 1 writer)
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	defer lock.Unlock()

	lock.Lock()
	b := balance
	balance = b + amount
}

// Reader (could be N readers)
func Balance(wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	defer lock.RUnlock()
	lock.RLock()
	b := balance
	fmt.Println("Current balance is", b)
}

func MainMutex() {
	var wg sync.WaitGroup
	var lock sync.RWMutex

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go Deposit(i, &wg, &lock)
		wg.Add(1)
		go Balance(&wg, &lock)

	}

	wg.Wait()
	fmt.Println("Finished with balance", balance)
}
