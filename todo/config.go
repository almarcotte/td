package todo

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Configuration struct {
	GlobalFile     string // Location of the global task list
	CurrentWorkDir string // Path where the program was executed from
	CurrentGlobal  bool   // If the current command is global or local
}

// NewConfiguration returns a struct containing settings we want to access across the entire program
func NewConfiguration() *Configuration {
	return &Configuration{
		GlobalFile:     globalFileLocation(),
		CurrentWorkDir: getCurrentWorkDir(),
	}
}

// getCurrentWorkDir returns the current directory, i.e. where the program was started from. This will be useful when
// figuring out if we're searching for a local (project) task list
func getCurrentWorkDir() string {
	this, err := os.Executable()

	if err != nil {
		log.Fatalf("Couldn't figure out where I am! Got: %s", err.Error())
	}

	return filepath.Dir(this)
}

// globalFileLocation returns the location of the global task list based on the user's home directory or an environment
// variable.
func globalFileLocation() (loc string) {
	loc = os.Getenv("TD_GLOBAL_LIST")

	if loc != "" {
		return
	}

	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(usr.HomeDir, ".td")
}
