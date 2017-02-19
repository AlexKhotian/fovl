package Networking

// SessionCommandType indicates behavior in session
type SessionCommandType uint32

// Type of session command
const (
	StartConnection SessionCommandType = 0
	StopConnection  SessionCommandType = 1
)

// JSON struct
type SessionCommand struct {
	CommandType SessionCommandType
}
