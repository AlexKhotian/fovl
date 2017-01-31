package TaskDistribution

// IJob Particular task to do
type IJob interface {
	PutTaskIn(newTask func())
	AbortTask()
	RunTask()
}

// Job handles a current task inside
type Job struct {
	_task        func()
	_syncChannel chan bool
}

// PutTaskIn stores task to value and wait what to do
func (_job *Job) PutTaskIn(newTask func()) {
	if _job._task != nil {
		_job._task = newTask
	}
}

// AbortTask ends current task
func (_job *Job) AbortTask() {
	_job._syncChannel <- false
	_job._task = nil
}

// RunTask executes task as go routine
func (_job *Job) RunTask() {
	go _job._task()
}
