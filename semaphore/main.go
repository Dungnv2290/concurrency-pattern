package main

import (
	"fmt"
	"time"
)

func main() {
	sem := NewSemaphore(3)
	doneC := make(chan bool, 1)
	totalProcess := 10
	for i := 1; i <= totalProcess; i++ {
		sem.Acquire()
		go func(v int) {
			defer sem.Release()
			longRunningProcess(v)
			if v == totalProcess {
				doneC <- true
			}
		}(i)
	}

	<-doneC
}

func longRunningProcess(taskID int) {
	fmt.Println(time.Now().Format("15:04:05"), "Running task with ID", taskID)
	time.Sleep(time.Second * 2)
}
