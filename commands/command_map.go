package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/config"
	"pokedexcli/utils"
	"strconv"
)

func CommandMap() error {
	res, err := http.Get(utils.LOCATION_API + "?offset=" + strconv.Itoa(config.Offset.Next) + "&limit=" + strconv.Itoa(utils.LIMIT))
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	apiRes := LocationResType{}
	err = json.Unmarshal(body, &apiRes)
	if err != nil {
		log.Fatal(err)
	}

	for _, location := range apiRes.Results {
		fmt.Printf("%s\n", location.Name)
	}
	fmt.Println()

	if config.Offset.Previous == 0 && config.Offset.Next == 0 {
		config.Offset.Next = utils.LIMIT
	} else {
		config.Offset.Next += utils.LIMIT
		config.Offset.Previous += utils.LIMIT
	}

	return nil
}
