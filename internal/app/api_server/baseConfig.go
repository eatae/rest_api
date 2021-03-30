package api_server


type baseConfig struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}


// NewApiServer...
// по соглашению используется как конструктур структуры
func NewBaseConfig() *baseConfig {
	return &baseConfig {
		// устанавливаем значения по умолчанию если они не заданы
		// в конфиге api_server.toml
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}