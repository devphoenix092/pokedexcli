package utils

type ApiType struct {
	Location     string
	LocationArea string
	Pokemon      string
}

var POKE_API = ApiType{
	Location:     "https://pokeapi.co/api/v2/location",
	LocationArea: "https://pokeapi.co/api/v2/location-area/",
	Pokemon:      "https://pokeapi.co/api/v2/pokemon/",
}

const LIMIT = 20
const CHANCE_LIMIT = 150
