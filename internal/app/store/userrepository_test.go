package store_test

import (
	"github.com/stretchr/testify/assert"
	"rest_api/internal/app/model"
	"rest_api/internal/app/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseUrl)
	defer teardown("users")

	email := "user@example.com"
	// проверяем, что user отсутсвует
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	// создаём user
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	// проверяем, что user присутствует
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}