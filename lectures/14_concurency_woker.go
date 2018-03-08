package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	fmt.Println("----workersDemo----------")
	workersDemo()

}

func workersDemo() {
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg, i)
	}

	for i := 0; i < tasks; i++ {
		tasksCh <- i
	}

	close(tasksCh)
}

func worker(tasksCh <-chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("worker:"+strconv.Itoa(workerID)+" processing task", task)
	}
}
