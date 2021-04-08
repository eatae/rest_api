package api_server

import "rest_api/internal/app/store"

type ServerConfig struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.StoreConfig
}

// New ServerConfig...
// по соглашению используется как конструктур структуры
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		// устанавливаем значения по умолчанию если они не заданы
		// в конфиге api_server.toml
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewStoreConfig(),
	}
}
