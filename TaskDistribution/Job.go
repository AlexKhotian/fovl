package TaskDistribution

// IJob Particular task to do
type IJob interface {
	PutTaskIn(newTask func(), channel chan bool)
	AbortTask()
	RunTask()
	GetID() uint64
	SetID(newID uint64)
}

// Job handles a current task inside
type Job struct {
	_task        func()
	_syncChannel chan bool
	_id          uint64
}

// JobFactory creates and init job
func JobFactory() IJob {
	job := new(Job)
	job._id = 0
	return job
}

// PutTaskIn stores task to value and wait what to do
func (_job *Job) PutTaskIn(newTask func(), channel chan bool) {
	if _job._task == nil {
		_job._task = newTask
		_job._syncChannel = channel
	}
}

// AbortTask ends current task
func (_job *Job) AbortTask() {
	if _job._syncChannel != nil {
		_job._syncChannel <- true
	}
	_job._task = nil
}

// RunTask executes task as go routine
func (_job *Job) RunTask() {
	if _job._task != nil {
		go _job._task()
	}
}

// GetID returns unique id of job
func (_job *Job) GetID() uint64 {
	return _job._id
}

// SetID gives new id to job
func (_job *Job) SetID(newID uint64) {
	_job._id = newID
}
