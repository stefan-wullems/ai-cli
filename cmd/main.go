package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided")
		os.Exit(1)
	}

	commands := []Command{
		&ChatCommand{},
	}
	subCommand := os.Args[1]
	for _, cmd := range commands {
		if cmd.Name() == subCommand {
			flagSet := flag.NewFlagSet(cmd.Name(), flag.ExitOnError)
			cmd.Init(flagSet)
			err := flagSet.Parse(os.Args[2:])
			if err != nil {
				fmt.Println("Error parsing flags:", err)
				os.Exit(1)
			}

			err = cmd.Execute(flagSet.Arg(0))
			if err != nil {
				fmt.Printf("Error executing %s command: %s\n", cmd.Name(), err)
				os.Exit(1)
			}
			return
		}
	}

	fmt.Printf("Unknown command: %s\n", subCommand)
	os.Exit(1)
}
