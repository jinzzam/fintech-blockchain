package command

import (
	"Blockchain/log"
	"strings"
)

// AddCommand adds a command for a global command.
func AddCommand(parent string, cmd Command) (err error) {
	if parent != "" {
		log.Debug("command.AddCommand: " + parent + " " + cmd.Name)
		c := globalCommands.FindFullCommand(parent)
		err = c.AddSubCommand(cmd)
	} else {
		log.Debug("command.AddCommand: " + cmd.Name)
		globalCommands = append(globalCommands, cmd)
	}

	return
}

// AddFlag adds a flag for a global command.
func AddFlag(parent string, flag Flag) error {
	c := globalCommands.FindFullCommand(parent)
	return c.AddFlag(flag)
}

// StartCommand executes a command for composited string.
func StartCommand(cmd string) error {
	cmd = strings.Replace(cmd, "\r", "", 1)
	log.Debug("command.StartCommand(" + cmd + ")")
	if cmd == "help" || cmd == "h" {
		globalCommands.Usage()
		return nil
	}
	return globalCommands.StartCommandWithFlags(cmd)
}
