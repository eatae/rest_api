package api_server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"rest_api/internal/app/store/teststore"
	"testing"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.NewStore())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
