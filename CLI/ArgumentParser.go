package CLI

import (
	"errors"
	"flag"
)

// ArgumentsRepresentation represents a command line args
type ArgumentsRepresentation struct {
	IsClient      bool
	SendFolder    string
	ReceiveFolder string
}

// ArgumentParser parses arguments from command line
type ArgumentParser interface {
	Parse() (ArgumentsRepresentation, error)
}

type argumentParserImpl struct {
}

// CreateArgumentParser creates instance of ArgumentParser
func CreateArgumentParser() ArgumentParser {
	return &argumentParserImpl{}
}

func (parser *argumentParserImpl) Parse() (ArgumentsRepresentation, error) {
	runType := flag.String("type", "server", "idecates if we are server / client")
	// TODO: think of better default path XD
	sendFolder := flag.String("sf", "/", "idecates what to send")
	receiveFolder := flag.String("rf", "/", "idecates where to receive")

	flag.Parse()

	isClient := false
	if *runType == "client" || *runType == "server" {
		isClient = (*runType == "client")
	} else {
		return ArgumentsRepresentation{}, errors.New("Failed to get runType")
	}
	return ArgumentsRepresentation{isClient, *sendFolder, *receiveFolder}, nil
}
