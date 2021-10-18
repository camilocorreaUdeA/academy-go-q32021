package workerspool

import (
	"fmt"
	"sync"
)

type WorkersPool struct {
	Jobs         []*Job
	workersCount int
	jobsQueue    chan *Job
	resultsQueue chan []string
	Results      [][]string
	wg           sync.WaitGroup
}

func NewWorkersPool(jobs []*Job, workers int) *WorkersPool {
	return &WorkersPool{
		Jobs:         jobs,
		workersCount: workers,
		jobsQueue:    make(chan *Job),
		resultsQueue: make(chan []string),
	}
}

func (wp *WorkersPool) Run() {
	for i := 1; i <= wp.workersCount; i++ {
		worker := NewWorker(wp.jobsQueue, wp.resultsQueue, i)
		worker.Start(&wp.wg)
	}

	for i := range wp.Jobs {
		wp.jobsQueue <- wp.Jobs[i]
		fmt.Println("Jobs:", i)
	}
	close(wp.jobsQueue)

	for i := 0; i < len(wp.Jobs); i++ {
		fmt.Println("results")
		wp.Results = append(wp.Results, <-wp.resultsQueue)
	}

	wp.wg.Wait()

	//close(wp.jobsQueue)
	//close(wp.resultsQueue)
	fmt.Println("here: end of worker pool Run")
}
