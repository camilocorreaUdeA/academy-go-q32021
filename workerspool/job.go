package workerspool

type Job struct {
	Error error
	Data  interface{}
	fn    func(interface{}) []string
}

// NewJob creates a new Job instance
func NewJob(fn func(interface{}) []string, data interface{}) *Job {
	return &Job{
		Data: data,
		fn:   fn,
	}
}

// Run runs the job function with job data as parameter
func (j *Job) Run() []string {
	return j.fn(j.Data)
}
