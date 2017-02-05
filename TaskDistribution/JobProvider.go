package TaskDistribution

import "container/list"

// IJobProvider manages active file transfer tasks to free workers in the pool
type IJobProvider interface {
	AddNewJob(newJob IJob)
	AbortJob(jobID uint64)
	GetJobsAmount() uint64
}

// JobProviderImpl impl for IJobProvider
type JobProviderImpl struct {
	_jobs        *list.List
	_jobsCounter uint64
}

// JobProviderFactory creates and inits job provider
func JobProviderFactory() IJobProvider {
	jobProvider := new(JobProviderImpl)
	jobProvider._jobs = list.New()
	jobProvider._jobsCounter = 1
	return jobProvider
}

// AddNewJob adds a new job to current job map with ID (uint64 TODO: Move to be a )
func (_jobProvide *JobProviderImpl) AddNewJob(newJob IJob) {
	newJob.SetID(_jobProvide._jobsCounter)
	_jobProvide._jobs.PushBack(newJob)
	_jobProvide._jobsCounter++
}

// AbortJob removes job from list
func (_jobProvide *JobProviderImpl) AbortJob(jobID uint64) {
	for element := _jobProvide._jobs.Front(); element != nil; element = element.Next() {
		job, _ := element.Value.(IJob)
		if job.GetID() == jobID {
			job.AbortTask()
			_jobProvide._jobs.Remove(element)
			return
		}
	}
}

// GetJobsAmount return how much currently running _jobs
func (_jobProvide *JobProviderImpl) GetJobsAmount() uint64 {
	return uint64(_jobProvide._jobs.Len())
}
