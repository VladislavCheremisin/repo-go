package main

import (
	"fmt"
	"sync"
)

var num int

func worker(i int, wait *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Printf("worker number %d is starting \n", i)
	mutex.Lock()
	num = num + 1
	mutex.Unlock()
	wait.Done()
	fmt.Printf("worker number %d has finished 5 operations\n", i)

}

func main() {
	var waitMain sync.WaitGroup
	var mutexMain sync.Mutex

	for i := 1; i <= 5; i++ {
		waitMain.Add(1)
		go worker(i, &waitMain, &mutexMain)
	}

	waitMain.Wait()

	fmt.Println("All workers have done work, value of workers i =", num)
}
