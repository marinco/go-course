package lecture3

import (
	"fmt"
	"sync"
)

var messages chan string
var myWaitGroup sync.WaitGroup

var wg sync.WaitGroup
var mux sync.Mutex
var sum int

func Main() {

	wg.Add(10000)
	var i int
	for i = 0; i < 10000; i++ {
		go foo(i)
	}
	wg.Wait()
	fmt.Println(sum)
}

func foo(i int) {
	mux.Lock()
	defer mux.Unlock()
	sum += 1
	//mux.Unlock()
	wg.Done()
}
