package TaskDistribution

// WorkerStatus represents what currently worker is doing
type WorkerStatus int

const (
	// HasNoJob means, that worker is open for new job
	HasNoJob WorkerStatus = 0
	// Busy - worker cannot accept new job
	Busy WorkerStatus = 1
)

// IWorker interface
type IWorker interface {
	AssignJob(newJob IJob)
	AbortJob()
	GetJobStatus() WorkerStatus
}

// Worker behavior and state
type Worker struct {
	_job    IJob
	_status WorkerStatus
}

// AssignJob gives job to user to care on
func (worker *Worker) AssignJob(newJob IJob) {
	if worker.GetJobStatus() == HasNoJob {
		worker._job = newJob
		worker._status = Busy
	}
}

// AbortJob stop currently runnung job
func (worker *Worker) AbortJob() {
	if worker.GetJobStatus() == Busy {
		worker._job.AbortTask()
		worker._job = nil
		worker._status = HasNoJob
	}
}

// GetJobStatus return a valid state of worker
func (worker *Worker) GetJobStatus() WorkerStatus {
	return worker._status
}
