package Networking

import (
	"encoding/json"
	"log"
	"net"
)

// IServerHandler interface for server handler functionality
// Start Up server
// Listen for incoming TCP(UDP) connection
// Start separate thread for handling a new session inside FileSessionManager
// Shutdown a connection
type IServerHandler interface {
	StartServer()
	Run()
	HandleNewConnection(conn *net.Conn)
	ShutdownServer()
}

// ServerHandler implementation for IServerHandler
type ServerHandler struct {
	_running bool
}

// ServerHanlderFactory creates and init new ServerHandler
func ServerHanlderFactory() IServerHandler {
	serverHandler := new(ServerHandler)
	serverHandler._running = true
	return serverHandler
}

// StartServer starts server
func (handler *ServerHandler) StartServer() {
	listener, listenErr := net.Listen("udp", ":9000")
	if listenErr != nil {
		log.Fatal(listenErr)
	}
	defer listener.Close()
	for {
		if handler._running == false {
			return
		}
		connection, connErr := listener.Accept()
		if connErr != nil {
			log.Fatal(connErr)
		}
		go handler.HandleNewConnection(&connection)
	}
}

// Run creates socket, listen for incoming connections
func (handler *ServerHandler) Run() {
	handler.StartServer()
}

// HandleNewConnection creates new thread for file transfer session
func (handler *ServerHandler) HandleNewConnection(conn *net.Conn) {
	log.Println("New incoming connection")
	var initData []byte
	(*conn).Read(initData)
	var cmd SessionCommand
	json.Unmarshal(initData, &cmd)
	if cmd.CommandType == StartConnectionRequest {
		// TODO(Khotian): Create repsonse with unique port id.
		// Spawn listener on this port.
	}
}

// ShutdownServer turn off server and clean up all threads
func (handler *ServerHandler) ShutdownServer() {
	handler._running = false
}
