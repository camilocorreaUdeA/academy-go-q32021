package workerspool

import "sync"

type WorkersPool struct {
	Jobs         []*Job
	workersCount int
	jobsQueue    chan *Job
	wg           sync.WaitGroup
}

func NewWorkersPool(jobs []*Job, workers int) *WorkersPool {
	return &WorkersPool{
		Jobs:         jobs,
		workersCount: workers,
		jobsQueue:    make(chan *Job),
	}
}

func (wp *WorkersPool) Run() {
	for i := 1; i <= wp.workersCount; i++ {
		worker := NewWorker(wp.jobsQueue, i)
		worker.Start(&wp.wg)
	}

	for i := range wp.Jobs {
		wp.jobsQueue <- wp.Jobs[i]
	}
	close(wp.jobsQueue)
	wp.wg.Wait()
}
