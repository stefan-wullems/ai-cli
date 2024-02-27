package main

import (
	"ai/chat"
	"flag"
	"fmt"
	"os"
)

type ChatCommand struct {
	outputFile string
	prompt     string
}

func (c *ChatCommand) Name() string {
	return "chat"
}

func (c *ChatCommand) Init(flags *flag.FlagSet) {
	flags.StringVar(&c.outputFile, "o", "", "Output file where the result will be saved")
}

func (c *ChatCommand) Execute(prompt string) error {
	if prompt == "" {
		return fmt.Errorf("prompt cannot be empty")
	}

	response, err := chat.Chat(c.prompt)
	if err != nil {
		return err
	}

	if c.outputFile != "" {
		// Write to the specified output file
		err := os.WriteFile(c.outputFile, []byte(response), 0644)
		if err != nil {
			return err
		}
	} else {
		// No output file specified, so just print to stdout
		fmt.Println(response)
	}

	return nil
}
