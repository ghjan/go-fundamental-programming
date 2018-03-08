package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

const (
	WORKERS    = 5
	SUBWORKERS = 3
	TASKS      = 20
	SUBTASKS   = 10
)

func main() {
	var wg sync.WaitGroup
	wg.Add(WORKERS)
	tasks := make(chan int)

	for i := 0; i < WORKERS; i++ {
		go worker2(tasks, &wg, i)
	}

	for i := 0; i < TASKS; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}

func worker2(tasks <-chan int, wg *sync.WaitGroup, workerID int) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		subtasks := make(chan int)
		for i := 0; i < SUBWORKERS; i++ {
			go subworker2(subtasks, workerID, i)
		}
		for i := 0; i < SUBTASKS; i++ {
			task1 := task*SUBTASKS + i
			subtasks <- task1
		}
		time.Sleep(2*time.Second)
		close(subtasks)
	}
}

func subworker2(subtasks chan int, workerID int, subworkerID int) {
	for {
		subtask, ok := <-subtasks
		if !ok {
			return
		}
		time.Sleep(time.Duration(subtask) * time.Millisecond)
		fmt.Println("workerID:"+strconv.Itoa(workerID)+", subworkerID:"+strconv.Itoa(subworkerID)+" processing subtask", subtask)
	}
}
