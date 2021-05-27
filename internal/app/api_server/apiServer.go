package api_server

import (
	"database/sql"
	"net/http"
	"rest_api/internal/app/store/sqlstore"
)

// Start ...
func Start(config *ServerConfig) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.NewStore(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
