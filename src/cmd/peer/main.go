package main

import (
	"blockchain/console"
	"blockchain/console/command"
	"blockchain/execute"
	"blockchain/log"
)

func main() {
	log.SetLogLevel(log.DebugLogLevel)
	RegisterCommand()
	console.Start()
}

func RegisterCommand() {
	execute.BlockchainCommands()
	execute.BlockCommands()
	execute.TransactionCommands()
	_ = command.AddCommand("", command.Command{
		Name:        "quit",
		ShortName:   "q",
		Description: "Exit the program",
		Commands:    nil,
		Flags:       nil,
		Run: func() error {
			console.GetContext().Quit()
			return nil
		},
	})
}
