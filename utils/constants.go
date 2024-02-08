package utils

type ApiType struct {
	Location     string
	LocationArea string
}

var POKE_API = ApiType{
	Location:     "https://pokeapi.co/api/v2/location",
	LocationArea: "https://pokeapi.co/api/v2/location-area/",
}

const LIMIT = 20
