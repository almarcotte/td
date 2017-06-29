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

func (add *AddCommand) Parse(conf *Configuration, args []string) error {
	return nil
}
