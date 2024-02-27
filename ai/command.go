package main

import "flag"

// Define a generic command interface
type Command interface {
	Name() string
	Init(flags *flag.FlagSet)
	Execute(arg string) error
}
