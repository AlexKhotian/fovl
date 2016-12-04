package Networking

// IServerHandler interface for server handler functionality
// Start Up server
// Listen for incoming TCP(UDP) connection
// Start separate thread for handling a new session inside FileSessionManager
// Shutdown a connection
type IServerHandler interface {
	StartServer()
	Run()
	HandleNewConnection()
	ShutdownServer()
}

// ServerHandler implementation for IServerHandler
type ServerHandler struct {
}

// StartServer starts server
func (handler *ServerHandler) StartServer() {

}

// Run creates socket, listen for incoming connections
func (handler *ServerHandler) Run() {

}

// HandleNewConnection creates new thread for file transfer session
func (handler *ServerHandler) HandleNewConnection() {

}

// ShutdownServer turn off server and clean up all threads
func (handler *ServerHandler) ShutdownServer() {

}
