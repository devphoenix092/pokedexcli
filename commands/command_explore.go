package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/cache"
	"pokedexcli/utils"
)

func CommandExplore(param ParamType) error {
	// Make api
	api := utils.POKE_API.LocationArea + "/" + param.Area

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
	apiRes := LocationAreaResType{}
	err := json.Unmarshal(val, &apiRes)
	if err != nil {
		return err
	}

	// Get location name
	fmt.Println("Exploring pastoria-city-area...")
	if len(apiRes.PokemonEncounters) != 0 {
		fmt.Println("Found Pokemon:")
		for _, location := range apiRes.PokemonEncounters {
			fmt.Printf("%s\n", location.Pokemon.Name)
		}
	} else {
		fmt.Println("Not Found Pokemon:")
	}

	fmt.Println()

	return nil
}
