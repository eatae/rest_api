package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq"
	"rest_api/internal/app/store"
)

// struct
type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

// NewStore
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return s.userRepository
}
