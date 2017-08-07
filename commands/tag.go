package commands

type TagCommand struct{}

func (c *TagCommand) Validate(configuration *Configuration) error {
	return nil
}

func (c *TagCommand) Parse(configuration *Configuration, args []string) error {
	return nil
}

func (c *TagCommand) Run(configuration *Configuration) error {
	return nil
}
