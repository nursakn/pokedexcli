package commands

func GetCommands() CommandsMap {
	return CommandsMap{
		"help": Command{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": Command{
			name: "exit",
			description: "Exit program",
			callback: commandExit,
		},
		"map": Command{
			name: "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback: commandMap,
		},
		"mapb": Command{
			name: "mapb",
			description: "Displays the names of 20 previous location areas in the Pokemon world. Each subsequent call to map should display the previous 20 locations, and so on.",
			callback: commandMapb,
		},
		"explore": Command{
			name: "explore",
			description: "Displays locations pokemons for passed name",
			callback: commandExplore,
		},
		"catch": Command{
			name: "catch",
			description: "Catches pokemon with given name",
			callback: commandCatch,
		},
		"inspect": Command{
			name: "inspect",
			description: "Inspects pokemon with given name if you have it in your pokedex",
			callback: commandInspect,
		},
		"pokedex": Command{
			name: "pokedex",
			description: "Shows what you have in your pokedex",
			callback: commandPokedex,
		},
	}
}
