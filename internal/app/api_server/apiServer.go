package api_server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// apiServer
type apiServer struct {
	config *baseConfig
	logger *logrus.Logger
	router *mux.Router
}

// NewApiServer
func NewApiServer(c *baseConfig) *apiServer {
	return &apiServer {
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}


// Start
// return error if not connect DB, or not work port...
func (s *apiServer) Start() error {
	// configureLogger
	if err := s.configureLogger(); err != nil {
		return err
	}
	// configureRouter
	s.configureRouter()
	// set Info
	s.logger.Info("Start api_server!")
	// Listen and Serve
	return http.ListenAndServe(s.config.BindAddr, s.router)
}



// configureLogger
func (s *apiServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}


// configureRouter
func (s *apiServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}





func (s *apiServer) handleHello() http.HandlerFunc {
	// ...
	// здесь можно устанавливать специфичные переменные для данного роута
	// ...

	return func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, "Hello world!")
	}
}















