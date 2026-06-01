package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationArea struct {
	Name string `json:"name"`
}

type locationAreaResponse struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

func commandMap(cfg *config) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	locationAreaRes := locationAreaResponse{}
	err = json.Unmarshal(body, &locationAreaRes)
	if err != nil {
		return err
	}

	if locationAreaRes.Next != nil {
		cfg.Next = *locationAreaRes.Next
	} else {
		cfg.Next = ""
	}

	if locationAreaRes.Previous != nil {
		cfg.Previous = *locationAreaRes.Previous
	} else {
		cfg.Previous = ""
	}

	for _, area := range locationAreaRes.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	url := cfg.Previous
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	locationAreaRes := locationAreaResponse{}
	err = json.Unmarshal(body, &locationAreaRes)
	if err != nil {
		return err
	}

	if locationAreaRes.Next != nil {
		cfg.Next = *locationAreaRes.Next
	} else {
		cfg.Next = ""
	}

	if locationAreaRes.Previous != nil {
		cfg.Previous = *locationAreaRes.Previous
	} else {
		cfg.Previous = ""
	}

	for _, area := range locationAreaRes.Results {
		fmt.Println(area.Name)
	}

	return nil
}
