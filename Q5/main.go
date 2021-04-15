//What will be printed when the code below is executed?
//And fix the issue to assure that `len(m)` is printed as 10.

package main

import (
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	/*wg.Add(1)
	go func() {
		for i := 0; i < N; i++ {
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	println(len(m))*/

	wg.Add(N)

	go func() {
		for i := 0; i < N; i++ {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}
	}()

	wg.Wait()
	println(len(m))
}
