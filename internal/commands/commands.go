package commands

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays list of 20 maps. Subsequent uses get next 20 maps.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays previous 20 maps, if not in first map page.",
			callback:    commandMapBack,
		},
	}
}
