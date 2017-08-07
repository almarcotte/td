package commands

type AddCommand struct {
	Due         string
	Tags        string
	Description string
}

func (add *AddCommand) Run(conf *Configuration) error {
	return nil
}

func (add *AddCommand) Validate(conf *Configuration) error {
	return nil
}

// Parse reads everything passed after the add command and tries to extract the description, due date and tags.
func (add *AddCommand) Parse(conf *Configuration, args []string) error {
	return nil
}
