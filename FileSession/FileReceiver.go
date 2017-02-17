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
	saveFiles()
}

// fileReceiver implements interface and perfomr save
type fileReceiver struct {
	_filePath string
}

// IFileReceiverFactory creates instance of receiver
func IFileReceiverFactory() IFileReceiver {
	fileReceiver := new(fileReceiver)
	return fileReceiver
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
	for err != nil && err != io.EOF {
		bytes, err := io.Copy(file, *conn)
	}
}

// saveFiles saves files to save path
func (receiver *fileReceiver) saveFiles() {

}
