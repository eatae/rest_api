package store

import (
	"database/sql"
	_ "github.com/lib/pq"
)

// struct
type Store struct {
	config         *StoreConfig
	db             *sql.DB
	userRepository *UserRepository
}

// NewStore
func NewStore(c *StoreConfig) *Store {
	return &Store{
		config: c,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DbUrl)
	if err != nil {
		return err
	}
	// check connection
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}

// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
