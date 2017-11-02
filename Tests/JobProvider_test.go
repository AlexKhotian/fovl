package Tests

import "testing"
import JobProvider "fovl/TaskDistribution"
import Job "fovl/TaskDistribution"

func TestCreateJobProvider(t *testing.T) {
	jobProvider := JobProvider.JobProviderFactory()
	if jobProvider == nil {
		t.Fail()
	}
}

func TestAddNewJob(t *testing.T) {
	jobProvider := JobProvider.JobProviderFactory()
	if jobProvider == nil {
		t.Error("No jobProvider")
	}
	job := new(Job.Job)
	if job == nil {
		t.Error("No job")
	}
	jobProvider.AddNewJob(job)
	if jobProvider.GetJobsAmount() != 1 {
		t.Error("Wrong count of jobs")
	}
}

func TestAddAndRemoveJob(t *testing.T) {
	jobProvider := JobProvider.JobProviderFactory()
	if jobProvider == nil {
		t.Error("No jobProvider")
	}
	job := new(Job.Job)
	if job == nil {
		t.Error("No job")
	}
	jobProvider.AddNewJob(job)
	if jobProvider.GetJobsAmount() != 1 {
		t.Error("Wrong count of jobs")
	}
	jobProvider.AbortJob(job.GetID())
	if jobProvider.GetJobsAmount() != 0 {
		t.Error("Wrong count of jobs")
	}
}
