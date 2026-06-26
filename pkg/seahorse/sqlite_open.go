//go:build !js

package seahorse

import (
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"
)

func openSQLite(dbPath string) (*sql.DB, error) {
	dsn := fmt.Sprintf("file:%s?_txlock=immediate", dbPath)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	return db, nil
}
