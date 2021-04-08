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
		DbUrl: "localhost:54321",
	}
}
