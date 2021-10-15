package workerspool

type Job struct {
	Error   error
	Data    interface{}
	fn      func(interface{}) []string
	MaxJobs int
}

func NewJob(fn func(interface{}) []string, data interface{}, max int) *Job {
	return &Job{
		Data:    data,
		fn:      fn,
		MaxJobs: max,
	}
}

func (j *Job) Run() []string {
	return j.fn(j.Data)
}
