package commands

type CommandType struct {
	Name        string
	Description string
	Callback    func(param ParamType) error
}

type LocationType struct {
	Name string `json:"name"`
}

type LocationResType struct {
	Results []LocationType `json:"results"`
}

type LocationAreaType struct {
	Pokemon LocationType `json:"pokemon"`
}

type LocationAreaResType struct {
	PokemonEncounters []LocationAreaType `json:"pokemon_encounters"`
}

type ParamType struct {
	Area        string
	PokemonName string
}

type PokemonResType struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetCommands() map[string]CommandType {
	return map[string]CommandType{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
		"map": {
			Name:        "map",
			Description: "Display the next 20 locations",
			Callback:    CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Display the previous 20 locations",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Display a list of all the Pokémon in a given area",
			Callback:    CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catching Pokemon",
			Callback:    CommandCatch,
		},
	}
}
