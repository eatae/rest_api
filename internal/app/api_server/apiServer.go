package api_server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"rest_api/internal/app/store"
)

// apiServer
type apiServer struct {
	config *ServerConfig
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// NewApiServer
func NewApiServer(c *ServerConfig) *apiServer {
	return &apiServer{
		config: c,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start
//
// return error if not connect DB, or not work port...
func (s *apiServer) Start() error {
	var err error = nil
	// configureLogger
	if err = s.configureLogger(); err != nil {
		return err
	}

	// configureRouter
	s.configureRouter()

	// configureStore
	if s.store, err = s.configureStore(); err != nil {
		return err
	}

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

// configureStore
func (s *apiServer) configureStore() (*store.Store, error) {
	pStore := store.NewStore(s.config.Store)
	return pStore, pStore.Open()
}

func (s *apiServer) handleHello() http.HandlerFunc {
	// ...
	// здесь можно устанавливать специфичные переменные для данного роута
	// ...

	return func(rw http.ResponseWriter, req *http.Request) {
		io.WriteString(rw, "Hello world!")
	}
}
