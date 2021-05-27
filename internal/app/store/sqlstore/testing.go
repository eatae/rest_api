package sqlstore

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"
)

// TestDB ...
func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	// помечаем функцию как тестовый хелпер
	t.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, `, `)))
			if err != nil {
				log.Fatal(err)
			}
		}
		db.Close()
	}
}
