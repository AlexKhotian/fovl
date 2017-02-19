package FileSession

// Define here json logic

// FileTransferCommandType defines behavior of filetrasnfer session
type FileTransferCommandType uint32

// Types of file transfer commands
const (
	Start    FileTransferCommandType = 0
	Finished FileTransferCommandType = 1
	Cancel   FileTransferCommandType = 2
	Pause    FileTransferCommandType = 3
	Resume   FileTransferCommandType = 4
)

// TransferCommand json struct
type TransferCommand struct {
	CommandType FileTransferCommandType
	DataChunk   []byte
}
