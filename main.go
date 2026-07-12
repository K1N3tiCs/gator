package main

import (
	"fmt"
	"github.com/K1N3tiCs/gator/internal/config"
	"log"
)

func main() {
	jsonConfig, err := config.Read()
	if err != nil {
		log.Fatal("failed reading the json config file")
	}
	fmt.Println(jsonConfig)

	if err := jsonConfig.SetUser("Penguuu"); err != nil {
		log.Fatal("failed to set user")
	}

	fmt.Println(jsonConfig)
}
