package cache

import (
	"fmt"
	"log"
	"time"
)

func MainCache() {
	cache := NewCache(GetFibonacci)

	fibo := []int{42, 40, 41, 42, 38}

	for _, n := range fibo {
		start := time.Now()
		value, err := cache.Get(n)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%d, %s, %d\n", n, time.Since(start), value)
	}
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
