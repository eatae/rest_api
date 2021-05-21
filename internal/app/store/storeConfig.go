package store

type StoreConfig struct {
	DbUrl string `toml:"db_url"`
}

// New StoreConfig
//
func NewStoreConfig() *StoreConfig {
	return &StoreConfig{
		// устанавливаем значения по умолчанию если они не заданы
		// в конфиге api_server.toml
		DbUrl: "postgresql://postgres:2222@localhost:54321/rest_api?sslmode=disable",
	}
}
