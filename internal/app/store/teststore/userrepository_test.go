package teststore_test

import (
	"github.com/stretchr/testify/assert"
	"rest_api/internal/app/model"
	"rest_api/internal/app/store"
	"rest_api/internal/app/store/teststore"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.NewStore()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.NewStore()
	email := "user@example.com"
	// проверяем, что user отсутсвует
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// создаём user
	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	// проверяем, что user присутствует
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
