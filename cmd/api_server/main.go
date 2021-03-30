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

//init
//Запуститься первой при вызове данного функционала
func init() {
	// можно передать путь до конфига с помощью параметра коммандной строки
	// по умолчанию значение configs/api_server.toml
	flag.StringVar(&configPath, "config-path", "configs/api_server.toml", "path to config file")
}


func main() {
	// получаем config-path из ком.стр или по умолчанию
	flag.Parse()
	// создаём структуру baseConfig
	config := api_server.NewBaseConfig()
	// заливаем данные из файла в структуру конфига
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := api_server.NewApiServer(config)
	// check start api_server
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
