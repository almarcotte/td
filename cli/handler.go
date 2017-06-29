package cli

import cmd "github.com/gnumast/td/commands"

const (
	// All the commands are listing here
	addCmd     = "a"
	addDue     = "due"
	addDueHelp = "The due date for this item. Accepts 'today', 'tomorrow', 'next week' or a datetime in the " +
		"'dd/mm/yyyy hh:ii:ss format"
	addTags     = "tag"
	addTagsHelp = "Comma-separated list of tags to apply to this item."

	usageCmd = ""
)

type Handler struct {
	Args []string
}

func NewHandler(args []string) *Handler {
	return &Handler{Args: args}
}

// Run tries to match the provided arguments with a command
func (h *Handler) Run(conf *cmd.Configuration) error {
	var c cmd.Command

	if h.Args[0] == usageCmd {
		c = cmd.UsageCommand{}
	}

	if err := c.Parse(conf, h.Args); err != nil {

	}

	if err := c.Validate(conf); err != nil {

	}

	if err := c.Run(conf); err != nil {

	}

	return nil
}
