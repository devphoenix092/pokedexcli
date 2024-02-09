package storage

type ConfigType struct {
	Next     int
	Previous int
}

var Offset ConfigType

type NameType struct {
	Name string `json:"name"`
}

type StatsType struct {
	BaseStat int      `json:"base_stat"`
	Stat     NameType `json:"stat"`
}

type TypesType struct {
	Type NameType `json:"type"`
}

type PokemonType struct {
	Name           string      `json:"name"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	Stats          []StatsType `json:"stats"`
	Types          []TypesType `json:"types"`
	BaseExperience int         `json:"base_experience"`
}

var PokemonList []PokemonType
