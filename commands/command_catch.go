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
	apiRes := PokemonResType{}
	err := json.Unmarshal(val, &apiRes)
	if err != nil {
		return err
	}

	// Get location name
	fmt.Println("Throwing a Pokeball at", param.PokemonName, "...")

	if rand.Intn(utils.CHANCE_LIMIT) > apiRes.BaseExperience {
		fmt.Println(param.PokemonName, "was caught!")
		pokemon := storage.PokemonType{
			Name: apiRes.Name,
		}
		storage.PokemonList = append(storage.PokemonList, pokemon)
	} else {
		fmt.Println(param.PokemonName, "escaped!")
	}

	fmt.Println()

	return nil
}
