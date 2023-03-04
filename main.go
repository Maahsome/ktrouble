package main

import (
	"os"

	"ktrouble/cmd"
)

func subCommands() (commandNames []string) {
	for _, command := range cmd.Commands() {
		commandNames = append(commandNames, append(command.Aliases, command.Name())...)
	}
	return
}

func setDefaultCommandIfNonePresent() {
	if len(os.Args) > 1 {
		potentialCommand := os.Args[1]
		for _, command := range subCommands() {
			if command == potentialCommand {
				return
			}
		}
	}
	os.Args = append([]string{os.Args[0], "launch"}, os.Args[1:]...)
}

func main() {
	setDefaultCommandIfNonePresent()
	cmd.Execute()
	// if err := cmd.Execute(); err != nil {
	// 	logrus.WithError(err).Error("Error executing command")
	// 	os.Exit(1)
	// }
}

// func main() {
// 	cmd.Execute()
// }
