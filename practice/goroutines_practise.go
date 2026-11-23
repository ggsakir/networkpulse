package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup, i int) {
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Println("Привет из горутины!", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	for i := range 3 {
		go sayHello(&wg, i)
	}
	wg.Wait()
	fmt.Println("Привет из main")
}
