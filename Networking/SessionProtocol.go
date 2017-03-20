package Networking

// SessionCommandType indicates behavior in session
type SessionCommandType uint32

// Type of session command
const (
	StartConnectionRequest  SessionCommandType = 0
	StartConnectionResponse SessionCommandType = 1
	StopConnection          SessionCommandType = 2
)

// JSON struct
type SessionCommand struct {
	CommandType SessionCommandType `json:"_commandType"`
	TCPPort     uint64             `json:"_tcpPort"`
}
