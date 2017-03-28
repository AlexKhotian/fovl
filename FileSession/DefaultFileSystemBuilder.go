package FileSession

import (
	"log"
	"os"
	"os/user"
)

const rootFileDir = "SmartFileTransport"
const mediaDir = "Media"
const photoDir = "Photo"
const docsDir = "Docs"
const codeDir = "SourceCode"

// CheckDefaultSubFileSystem check if SmartFileTransport file system was initialized
func CheckDefaultSubFileSystem() {
	if _, err := os.Stat(rootFileDir); os.IsNotExist(err) {
		usrToken, err := user.Current()
		if err != nil {
			log.Fatal("Failed to get current user")
			return
		}
		homeDir := usrToken.HomeDir
		if _, err := os.Stat(homeDir); err == nil {
			log.Fatal("Failed to get current user home dir")
			return
		}
	} else {
		log.Fatal("Directory already exists")
	}
}

func createDefaultFileSystem(userHomeDir string) {
	if _, err := os.Create(mediaDir); err != nil {
		log.Fatal("Failed to create Media folder")
	}

	if _, err := os.Create(photoDir); err != nil {
		log.Fatal("Failed to create Photo folder")
	}

	if _, err := os.Create(docsDir); err != nil {
		log.Fatal("Failed to create Docs folder")
	}

	if _, err := os.Create(codeDir); err != nil {
		log.Fatal("Failed to create Source code folder")
	}
}
