package teststore

import (
	"rest_api/internal/app/model"
	"rest_api/internal/app/store"
)

// struct
type Store struct {
	userRepository *UserRepository
}

// NewStore
func NewStore() *Store {
	return &Store{}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}
