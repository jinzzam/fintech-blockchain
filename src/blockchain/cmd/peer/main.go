package main

func main() {
	log.SetLogLevel(log.DebugLogLevel)
	RegisterCommand()
	console.Start()
}

func RegisterCommand() {
	excute.BlockchainCommands()
	excute.BlockCommands()
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
