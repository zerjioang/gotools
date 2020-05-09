package worker

import (
	"github.com/zerjioang/gotools/lib/logger"
)

// A buffered channel that we can send work requests on.
var JobQueue chan JobInterface

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan JobInterface
	JobChannel chan JobInterface
	quit       chan bool
}

const (
	maxWorker = 4
)

func init() {
	dispatcher := NewDispatcher(maxWorker)
	dispatcher.Run()
}

func NewWorker(workerPool chan chan JobInterface) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan JobInterface),
		quit:       make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			// register the current worker into the worker queue.
			w.WorkerPool <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				// we have received a work request.
				if err := job.Run(); err != nil {
					logger.Error("error executing the job: ", err.Error())
				}

			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func SubmitJob(job JobInterface) {
	// let's submit a job to the work onto the queue.
	JobQueue <- job
}
