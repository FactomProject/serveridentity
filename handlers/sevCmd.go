package handlers

type sevCmd struct {
	execFunc    func([]string)
	helpMsg     string
	description string
}

func (c *sevCmd) Execute(args []string) {
	c.execFunc(args)
}
