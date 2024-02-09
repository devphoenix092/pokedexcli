package commands

import (
	"fmt"
	"pokedexcli/storage"
)

func CommandInspect(param ParamType) error {
	pokemon := storage.PokemonType{}
	for _, item := range storage.PokemonList {
		if item.Name == param.PokemonName {
			pokemon = item
		}
	}

	if pokemon.Name != "" {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)
		fmt.Println("Stats:")
		for _, item := range pokemon.Stats {
			fmt.Println("-", item.Stat.Name)
		}
		fmt.Println("Types:")
		for _, item := range pokemon.Types {
			fmt.Println("-", item.Type.Name)
		}
	} else {
		fmt.Println("You haven't caught this Pokémon yet.")
	}

	fmt.Println()

	return nil
}
