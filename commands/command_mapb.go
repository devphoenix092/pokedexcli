package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/cache"
	"pokedexcli/config"
	"pokedexcli/utils"
	"strconv"
)

func CommandMapb() error {
	// Change config
	if config.Offset.Previous != 0 {
		config.Offset.Next -= utils.LIMIT
		config.Offset.Previous -= utils.LIMIT
	}

	// Make api
	api := utils.LOCATION_API + "?offset=" + strconv.Itoa(config.Offset.Previous) + "&limit=" + strconv.Itoa(utils.LIMIT)

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
	apiRes := LocationResType{}
	err := json.Unmarshal(val, &apiRes)
	if err != nil {
		return err
	}

	// Get location name
	for _, location := range apiRes.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()

	return nil
}
