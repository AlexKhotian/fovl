package TaskDistribution

// IJobProvider manages active file transfer tasks to free workers in the pool
type IJobProvider interface {
    AddNewJob()
    AbortJob()
    GetJobID()
    
}

// JobProviderImpl impl for IJobProvider
type JobProviderImpl struct {
}
