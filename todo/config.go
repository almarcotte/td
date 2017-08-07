package todo

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type Configuration struct {
	GlobalFile string
}

func NewConfiguration() *Configuration {
	return &Configuration{
		GlobalFile: globalFileLocation(),
	}
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
