package commands

const (
	// All the commands are listing here
	addCmd      = "a"
	completeCmd = "c"
	progressCmd = "p"
	revertCmd   = "r"
	tagCmd      = "t"
	listCmd     = "ls"
)

type Handler struct {
	Args []string
}

func NewHandler(args []string) *Handler {
	return &Handler{Args: args}
}

// Run tries to match the provided arguments with a command
func (h *Handler) Run(conf *Configuration) error {
	var c Command

	c = parseArgsForCommand(h.Args[0])

	if err := c.Parse(conf, h.Args); err != nil {

	}

	if err := c.Validate(conf); err != nil {

	}

	if err := c.Run(conf); err != nil {

	}

	return nil
}

func parseArgsForCommand(arg string) Command {
	// This feels a bit messy, maybe there's a better way to do this..?
	switch arg {
	case addCmd:
		return &AddCommand{}
	case completeCmd:
		return &StatusCommand{Flag: completeCmd}
	case progressCmd:
		return &StatusCommand{Flag: progressCmd}
	case revertCmd:
		return &StatusCommand{Flag: revertCmd}
	case tagCmd:
		return &TagCommand{}
	case listCmd:
		return &ListCommand{}
	}

	return &UsageCommand{}
}
