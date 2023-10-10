package main

import "log"

type WorkerPool interface {
	Run()
	Add(task func())
}

type workerPool struct {
	maxWorker  int
	queuedTask chan func()
}

func NewWorkerPool(maxWorker int) WorkerPool {
	wp := &workerPool{
		maxWorker:  maxWorker,
		queuedTask: make(chan func()),
	}

	return wp
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		wID := i + 1
		log.Printf("[WorkerPool] Worker %d has been spawned", wID)

		go func(workerID int) {
			for task := range wp.queuedTask {
				log.Printf("[WorkerPool] Worker %d start processing task", workerID)
				task()
				log.Printf("[WorkerPool] Worker %d finish processing task", workerID)
			}
		}(wID)
	}
}

func (wp *workerPool) Add(task func()) {
	wp.queuedTask <- task
}
