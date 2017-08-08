package commands

import (
	"errors"
	"fmt"
	"github.com/gnumast/td/cli"
	"strings"
)

const (
	addCmd      = "a"
	completeCmd = "c"
	progressCmd = "p"
	revertCmd   = "r"
	tagCmd      = "t"
	listCmd     = "ls"
	versionCmd  = "version"
)

type Handler struct {
	Args []string
}

// NewHandler returns a new command handler
func NewHandler(args []string) *Handler {
	return &Handler{Args: args[1:]}
}

// Run tries to match the provided arguments with a command then asks it to parse / validate the rest of the arguments.
// Finally the command is executed.
func (h *Handler) Run(app *cli.Application) (err error) {
	command, err := h.ParseForCommand()

	if err != nil {
		app.CliOutput.Error(err)
	}

	// Parsing error probably means missing arguments
	if err = command.Parse(app, h.Args); err != nil {
		command.Help(app)

		return
	}

	// Validation error means the arguments were received don't make sense, such as removing a task that doesn't exist
	if err = command.Validate(app); err != nil {
		command.Help(app)

		return
	}

	// Something went wrong while executing the command, hopefully this doesn't happen
	if err = command.Run(app); err != nil {
		app.CliOutput.Error(err)
	}

	return
}

// parseArgsForCommand receives the first argument passed via the command line and returns the appropriate command
func (h *Handler) ParseForCommand() (cmd Command, err error) {
	arg := ""

	if len(h.Args) > 0 {
		arg = h.Args[0]
	}

	switch strings.ToLower(arg) {
	case addCmd:
		cmd = &AddCommand{}
	case completeCmd:
		cmd = &StatusCommand{Flag: completeCmd}
	case progressCmd:
		cmd = &StatusCommand{Flag: progressCmd}
	case revertCmd:
		cmd = &StatusCommand{Flag: revertCmd}
	case tagCmd:
		cmd = &TagCommand{}
	case listCmd:
		cmd = &ListCommand{}
	case versionCmd:
		cmd = &VersionCommand{}
	default:
		cmd = &UsageCommand{}
		if arg != "" {
			err = errors.New(fmt.Sprintf("No such command `%s`", arg))
		}
	}

	return
}
