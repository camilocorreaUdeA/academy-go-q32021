package workerspool

import (
	"fmt"
	"log"
	"sync"
)

type Worker struct {
	ID       int
	jobChan  chan *Job
	resChan  chan []string
	doneJobs int
}

func NewWorker(jobsChannel chan *Job, resChannel chan []string, workerID int) *Worker {
	return &Worker{
		ID:      workerID,
		jobChan: jobsChannel,
		resChan: resChannel,
	}
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for job := range w.jobChan {
			fmt.Println("jobchan")
			res := job.Run()
			if res != nil && w.doneJobs < job.MaxJobs {
				w.doneJobs += 1
				log.Printf("Jobs completed so far %d by worker %d", w.doneJobs, w.ID)
				log.Println("result added to queue")
				w.resChan <- res
			}
		}
		fmt.Println("Done!!")
	}()
}
