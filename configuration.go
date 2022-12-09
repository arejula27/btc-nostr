package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Bitcoin BitcoinConf
	Nostr   NostrConf
}

// init func runs before main, it will load the configuration
func loadConfiguration() (*Config, error) {
	data, err := ioutil.ReadFile("./conf")
	if err != nil {
		fmt.Println(err)
	}

	config := Config{}
	err = toml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println("Configuration loaded")
	return &config, nil

}
