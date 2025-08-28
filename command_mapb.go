package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"encoding/json"
)

func commandMapb(c *config) error {

	if c.Previous == "" {
		fmt.Println("You're already on the first page.")
		return nil
	}

	res, err := http.Get(c.Previous)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299  {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}


	err = json.Unmarshal(body, c)
	if err != nil {
		fmt.Println(err)
	}
	
	type LocationAreas struct {
		Results []struct {
			Name string `json:"name"`
		}
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}

	var locationList LocationAreas
	err = json.Unmarshal(body, &locationList)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range locationList.Results {
		fmt.Println(loc.Name)
	}

	c.Next = locationList.Next
	c.Previous = locationList.Previous

	return nil
}