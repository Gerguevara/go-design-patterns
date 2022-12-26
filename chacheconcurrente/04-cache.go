package cacheconcurrente

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func MainCache() {
	cache := NewCache(GetFibonacci)
	fibo := []int{42, 40, 41, 42, 38}
	var wg sync.WaitGroup
	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, err := cache.Get(index)
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("%d, %s, %d\n", index, time.Since(start), value)
		}(n)

	}
	wg.Wait()
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// el retorno de interface practicemente dice qeu retorna un generico
// en este caso retorna la funcion (O sea el resultado y error)
func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}
