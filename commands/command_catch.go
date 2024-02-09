package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"pokedexcli/cache"
	"pokedexcli/storage"
	"pokedexcli/utils"
)

func CommandCatch(param ParamType) error {
	// Make api
	api := utils.POKE_API.Pokemon + "/" + param.PokemonName

	// Check cache
	val, isExists := cache.Get(api)
	if !isExists {
		// Get the data from api
		res, err := http.Get(api)
		if err != nil {
			return err
		}

		val, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		res.Body.Close()

		// Add new cache
		cache.Add(api, val)
	}

	// json decode
	apiRes := storage.PokemonType{}
	err := json.Unmarshal(val, &apiRes)
	if err != nil {
		return err
	}

	isCatched := false
	for _, item := range storage.PokemonList {
		if item.Name == param.PokemonName {
			isCatched = true
		}
	}

	if !isCatched {
		fmt.Println("Throwing a Pokeball at", param.PokemonName, "...")

		chanceValue := rand.Intn(utils.CHANCE_LIMIT)
		if chanceValue > apiRes.BaseExperience {
			fmt.Println(param.PokemonName, "was caught!")
			fmt.Println(param.PokemonName, "You may now inspect it with the inspect command.")
			pokemon := storage.PokemonType{
				Name:   apiRes.Name,
				Height: apiRes.Height,
				Weight: apiRes.Weight,
				Stats:  apiRes.Stats,
				Types:  apiRes.Types,
			}

			storage.PokemonList = append(storage.PokemonList, pokemon)
		} else {
			fmt.Println(param.PokemonName, "escaped!")
		}
	} else {
		fmt.Println("You have already caught this pokemon.")
	}

	fmt.Println()

	return nil
}
