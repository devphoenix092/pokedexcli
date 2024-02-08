package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedexcli/cache"
	"pokedexcli/storage"
	"pokedexcli/utils"
	"strconv"
)

func CommandMap(param ParamType) error {
	// Make api
	api := utils.POKE_API.Location + "?offset=" + strconv.Itoa(storage.Offset.Next) + "&limit=" + strconv.Itoa(utils.LIMIT)

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

	// Change config
	if storage.Offset.Previous == 0 && storage.Offset.Next == 0 {
		storage.Offset.Next = utils.LIMIT
	} else {
		storage.Offset.Next += utils.LIMIT
		storage.Offset.Previous += utils.LIMIT
	}

	return nil
}
