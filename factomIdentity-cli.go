package main

import (
	"flag"
	"github.com/FactomProject/cli"
	"github.com/FactomProject/serveridentity/handlers"
)

// Testing CLI
func main() {
	flag.Parse()
	args := flag.Args()

	c := cli.New()

	c.Handle("help", handlers.Help)
	c.Handle("start", handlers.Start)
	c.Handle("mhash", handlers.NewMHash)
	c.Handle("newkey", handlers.NewKey)

	c.HandleDefault(handlers.Help)
	c.Execute(args)
}
