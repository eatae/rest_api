package api_server

// ServerConfig
//
// Cтруктура описываящая конфиг сервера
// который мы передаём в конструктор apiServer
type ServerConfig struct {
	// порт сервера
	BindAddr string `toml:"bind_addr"`
	// базовый уровень логирования (читаем из toml файла)
	LogLevel    string `toml:"log_level"`
	DatabaseUrl string `toml:"database_url"`
}

// New ServerConfig...
//
// по соглашению используется как конструктур
func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		// устанавливаем значения по умолчанию если они не заданы
		// в конфиге api_server.toml
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
