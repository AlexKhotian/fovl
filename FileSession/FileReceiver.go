package FileSession

import (
	"net"
)

// IFileReceiver receiver file and folders from network.
type IFileReceiver interface {
	SetSavePath(filePath *string)
	ReceiveFileFromConnection(conn net.Conn)
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
func (receiver *fileReceiver) ReceiveFileFromConnection(conn net.Conn) {

}

// saveFiles saves files to save path
func (receiver *fileReceiver) saveFiles() {

}
