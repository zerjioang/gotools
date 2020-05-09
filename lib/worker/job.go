package worker

type JobInterface interface {
	Run() error
}

// Job represents the job to be run
type Job struct {
	JobInterface
	Payload []byte
}

func (j *Job) Run() error {
	return nil
}
