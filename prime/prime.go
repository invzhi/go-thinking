package prime

import (
	"fmt"
	"math"
)

func generate(ch chan<- int, n int) {
	for i := 2; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)
}

func GetSlices(n int) ([]int, error) {
	if n < 2 {
		return nil, fmt.Errorf("cannot get primes which <= %d", n)
	}
	// x/ln(x) is primes' amount
	primes := make([]int, 0, n/int(math.Log(float64(n))))

	ch := make(chan int)
	go generate(ch, n)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		primes = append(primes, prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	return primes, nil
}
