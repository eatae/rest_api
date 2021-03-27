package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"rest_api/internal/app/api_server"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/api_server.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := api_server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := api_server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
