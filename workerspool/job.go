package workerspool

type Job struct {
	Error error
	Data  interface{}
	fn    func(interface{})
}

func NewJob(fn func(interface{}), data interface{}) *Job {
	return &Job{
		Data: data,
		fn:   fn,
	}
}

func (j *Job) Run() {
	j.fn(j.Data)
}
