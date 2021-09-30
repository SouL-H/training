package main

import (
	"fmt"
	"sync"
)

type RaceTest struct {
	Val int
}

func main() {
	raceTest := &RaceTest{}
	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go increment(raceTest, wg, lock)
	}

	wg.Wait()

	fmt.Println(raceTest)
}

//Race-Condition fixed
func increment(rt *RaceTest, wg *sync.WaitGroup, lk *sync.Mutex) {
	lk.Lock()
	rt.Val += 1
	lk.Unlock()
	wg.Done()
}
