package main

import (
	"log"
	"runtime"
	"time"
)

func main() {
	waitC := make(chan bool)
	go func() {
		for {
			log.Printf("[main] Total current goroutine: %d", runtime.NumGoroutine())
			time.Sleep(1 * time.Second)
		}
	}()

	totalWorker := 3
	wp := NewWorkerPool(totalWorker)
	wp.Run()

	type result struct {
		id    int
		value int
	}

	totalTask := 5
	resultC := make(chan result, totalTask)

	for i := 0; i < totalTask; i++ {
		id := i + 1
		wp.Add(func() {
			log.Printf("[main] Starting task %d", id)
			time.Sleep(5 * time.Second)
			resultC <- result{id, id * 2}
		})
	}

	for i := 0; i < totalTask; i++ {
		res := <-resultC
		log.Printf("[main] Task %d has been finished with result %d:", res.id, res.value)
	}

	<-waitC
}
