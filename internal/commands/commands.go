package commands

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *Config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays list of 20 maps. Subsequent uses get next 20 maps.",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous 20 maps, if not in first map page.",
			Callback:    commandMapBack,
		},
	}
}
