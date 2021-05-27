package sqlstore_test

import (
	"os"
	"testing"
)

var databaseUrl string

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "postgres://postgres:2222@localhost:54321/test_rest_api?sslmode=disable"
	}
	os.Exit(m.Run())
}
