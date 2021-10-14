package workerspool

import (
	"fmt"
	"sync"
)

type Worker struct {
	ID      int
	jobChan chan *Job
}

func NewWorker(jobsChannel chan *Job, workerID int) *Worker {
	return &Worker{
		ID:      workerID,
		jobChan: jobsChannel,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for job := range w.jobChan {
			fmt.Print("worker ", w.ID, " ")
			job.Run()
		}
	}()
}
