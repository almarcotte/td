package commands

type ListCommand struct{}

func (c *ListCommand) Validate(configuration *Configuration) error {
	return nil
}

func (c *ListCommand) Parse(configuration *Configuration, args []string) error {
	return nil
}

func (c *ListCommand) Run(configuration *Configuration) error {
	return nil
}
