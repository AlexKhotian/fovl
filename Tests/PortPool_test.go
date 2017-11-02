package Tests

import (
	"fovl/Networking"
	"testing"
)

func TestGetNewPortIDFromPool(t *testing.T) {
	pool := Networking.PortPoolFactory()
	if pool == nil {
		t.Log("Failed to create pool")
		t.FailNow()
	}
	portID := pool.GetNewPortID()
	if portID == 0 {
		t.Log("Failed to get portID from pool")
		t.FailNow()
	}

	if !pool.IsPortIDInUse(portID) {
		t.Log("Failed: portID is not in use, but it has to be")
		t.FailNow()
	}
	t.Log("GetNewPortIDFromPoolTest Passed")
}

func TestReleasePortIDFromUsage(t *testing.T) {
	pool := Networking.PortPoolFactory()
	if pool == nil {
		t.Log("Failed to create pool")
		t.FailNow()
	}
	portID := pool.GetNewPortID()

	pool.ReleasePortID(portID)
	if pool.IsPortIDInUse(portID) {
		t.Log("Failed: portID is in use, but it has not to be")
		t.FailNow()
	}
	t.Logf("TestReleasePortIDFromUsage Passed")
}
