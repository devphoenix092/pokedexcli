package storage

type ConfigType struct {
	Next     int
	Previous int
}

var Offset ConfigType

type PokemonType struct {
	Name string
}

var PokemonList []PokemonType
