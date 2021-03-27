package api_server

// APIServer
type APIServer struct {
	config *Config
}

// New
func New(conf *Config) *APIServer {
	return &APIServer{
		config: conf,
	}
}

// Start
func (s APIServer) Start() error {
	return nil
}
