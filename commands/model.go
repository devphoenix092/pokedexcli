package commands

type CommandType struct {
	Name        string
	Description string
	Callback    func() error
}

type LocationType struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationResType struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationType `json:"results"`
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
	}
}
