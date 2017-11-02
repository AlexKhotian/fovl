package Tests

import "testing"
import "fovl/TaskDistribution"
import "fmt"

func TestCreateJob(t *testing.T) {
	job := TaskDistribution.JobFactory()
	if job == nil {
		t.Error("No job")
	}
}

func TestCreateJobAndAssignTask(t *testing.T) {
	job := TaskDistribution.JobFactory()
	if job == nil {
		t.Error("No job")
	}
	c := make(chan bool)
	task := func() {
		for i := 0; i != 5; i++ {
		}
		c <- true
	}
	job.PutTaskIn(task, c)
	job.RunTask()
	select {
	case <-c:
		t.Log("Got finish")
	}
}

func TestCreateJobAndAssignTaskAndRemoveTask(t *testing.T) {
	job := TaskDistribution.JobFactory()
	if job == nil {
		t.Error("No job")
	}
	c := make(chan bool)
	task := func() {
		for i := 0; i != 10000000; i++ {
			fmt.Println("Test!")
			if true == <-c {
				return
			}
		}
		c <- true
	}
	job.PutTaskIn(task, c)
	job.RunTask()
	job.AbortTask()
	/*select {
	case <-c:
		t.Log("Got finish")
	}*/
}
