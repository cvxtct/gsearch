package main

import (
    "encoding/json"
    "os"
    "fmt"
)

func Configuration() Config {
	file, _ := os.Open("../config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	// TODO Validate config
	return configuration
}