package Networking

// IPortPool handles availabel and in-use ports
type IPortPool interface {
	GetNewPortID() uint32
	ReleasePortID(portID uint32)
	IsPortIDInUse(portID uint32) bool
}

type portPool struct {
	_portPool map[uint32]bool
}

// PortPoolFactory creates an instance of IPortPool
func PortPoolFactory() IPortPool {
	newPortPool := new(portPool)
	newPortPool._portPool = make(map[uint32]bool)
	startPortID := 27000
	endPortID := 28000
	for i := startPortID; i <= endPortID; i++ {
		newPortPool._portPool[uint32(i)] = false
	}
	return newPortPool
}

// GetNewPortId looks up for free port in pool
func (pool *portPool) GetNewPortID() uint32 {
	for portID, isInUse := range pool._portPool {
		if isInUse == false {
			pool._portPool[portID] = true
			return portID
		}
	}
	return 27000
}

func (pool *portPool) ReleasePortID(portID uint32) {
	pool._portPool[portID] = false
}

func (pool *portPool) IsPortIDInUse(portID uint32) bool {
	return pool._portPool[portID]
}
