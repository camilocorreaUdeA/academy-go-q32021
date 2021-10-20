package workerspool

import (
	"log"
	"sync"
)

type WorkersPool struct {
	Jobs         []*Job
	workersCount int
	jobsQueue    chan *Job
	resultsQueue chan []string
	Results      [][]string
	maxJobs      int
	wg           sync.WaitGroup
}

// NewWorkersPool returns an instance of the worker pool
func NewWorkersPool(jobs []*Job, workers int, max int) *WorkersPool {
	return &WorkersPool{
		Jobs:         jobs,
		workersCount: workers,
		jobsQueue:    make(chan *Job),
		resultsQueue: make(chan []string, len(jobs)),
		maxJobs:      max,
	}
}

// Run executes all jobs by passing them to workers, reads results and returns to the service
func (wp *WorkersPool) Run() [][]string {
	for i := 1; i <= wp.workersCount; i++ {
		worker := NewWorker(wp.jobsQueue, wp.resultsQueue, i, wp.maxJobs)
		worker.Start(&wp.wg)
	}

	for i := range wp.Jobs {
		wp.jobsQueue <- wp.Jobs[i]
	}
	close(wp.jobsQueue)

	wp.wg.Wait()
	log.Println("All workers finished their jobs")

	close(wp.resultsQueue)
	for res := range wp.resultsQueue {
		wp.Results = append(wp.Results, res)
	}

	return wp.Results
}
