package commands

import (
	"fmt"
	"pokedexcli/storage"
)

func CommandPokedex(param ParamType) error {
	if len(storage.PokemonList) == 0 {
		fmt.Println("You haven't caught any Pokémons yet.")
	} else {
		fmt.Println("Your Pokedex:")
		for _, item := range storage.PokemonList {
			fmt.Println("-", item.Name)
		}
	}

	fmt.Println()

	return nil
}
