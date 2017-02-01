package TaskDistribution

// IJobProvider manages active file transfer tasks to free workers in the pool
type IJobProvider interface {
	AddNewJob(newJob IJob)
	AbortJob(jobID uint64)
	GetJobID() uint64
	GetJobsAmount() uint64
}

// JobProviderImpl impl for IJobProvider
type JobProviderImpl struct {
	_jobs map[uint64]IJob
}

/*
// JobProviderFactory creates and inits job provider
func JobProviderFactory() IJobProvider {

}

// AddNewJob adds a new job to current job map with ID (uint64 TODO: Move to be a )
func (_jobProvide *JobProviderImpl) AddNewJob(newJob IJob) {
	UUIDGenerator

	_jobProvide._jobs
}*/
