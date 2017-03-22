package FileSession

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// IFileReceiver receiver file and folders from network.
type IFileReceiver interface {
	SetSavePath(filePath *string)
	ReceiveFileFromConnection(conn *net.Conn)

	initFileReceiver() bool
}

// fileReceiver implements interface and perfomr save
type fileReceiver struct {
	_filePath          string
	_listener          net.Listener
	_connection        net.Conn
	_transmissionEnded chan bool
}

// IFileReceiverFactory creates instance of receiver
func IFileReceiverFactory() IFileReceiver {
	fileReceiver := new(fileReceiver)
	return fileReceiver
}

func (receiver *fileReceiver) initFileReceiver() bool {
	listener, err := net.Listen("tcp", "localhost:27111")
	if err != nil {
		log.Fatal(err)
		return false
	}
	receiver._listener = listener
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
		return false
	}
	receiver._connection = conn
	receiver.ReceiveFileFromConnection(&conn)
	return true
}

// SetSavePath set path where to save new saveFiles
func (receiver *fileReceiver) SetSavePath(filePath *string) {
	receiver._filePath = *filePath
}

// ReceiveFileFromConnection starts to receive file
func (receiver *fileReceiver) ReceiveFileFromConnection(conn *net.Conn) {
	file, err := os.Open(strings.TrimSpace(receiver._filePath))
	if err != nil {
		log.Fatal("Failed to create file")
	}
	defer file.Close()

	for {
		io.CopyN(file, *conn, 1024)
	}
}
