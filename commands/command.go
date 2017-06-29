package commands

// Command represents a supported action that can parse command line arguments, validate them and run.
type Command interface {
	Run(*Configuration) error
	Validate(*Configuration) error
	Parse(*Configuration, []string) error
}
