package handlers

type fctCmd struct {
	execFunc    func([]string)
	helpMsg     string
	description string
}

func (c *fctCmd) Execute(args []string) {
	c.execFunc(args)
}
