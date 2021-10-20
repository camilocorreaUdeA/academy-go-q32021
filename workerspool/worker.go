package workerspool

import (
	"log"
	"sync"
)

type Worker struct {
	ID       int
	jobChan  chan *Job
	resChan  chan []string
	doneJobs int
	maxJobs  int
}

// NewWorker creates a new worker that will run in a new goroutine
func NewWorker(jobsChannel chan *Job, resChannel chan []string, workerID int, max int) *Worker {
	return &Worker{
		ID:      workerID,
		jobChan: jobsChannel,
		resChan: resChannel,
		maxJobs: max,
	}
}

// Start adds a new goroutine for each worker to execute a job
func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for job := range w.jobChan {
			res := job.Run()
			if res != nil && w.doneJobs < w.maxJobs {
				w.doneJobs += 1
				log.Printf("Jobs completed so far %d by worker %d", w.doneJobs, w.ID)
				log.Println("result added to queue")
				w.resChan <- res
			}
		}
	}()
}
