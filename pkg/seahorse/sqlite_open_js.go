//go:build js

package seahorse

import (
	"database/sql"
	"fmt"

	"github.com/ncruces/go-sqlite3"
	"github.com/ncruces/go-sqlite3/driver"
)

func openSQLite(dbPath string) (*sql.DB, error) {
	dsn := fmt.Sprintf("file:%s?_txlock=immediate&vfs=idb", dbPath)
	db, err := driver.Open(dsn, func(c *sqlite3.Conn) error {
		for _, pragma := range []string{
			"PRAGMA journal_mode = WAL",
			"PRAGMA busy_timeout = 5000",
			"PRAGMA synchronous = NORMAL",
		} {
			if err := c.Exec(pragma); err != nil {
				return fmt.Errorf("set pragma: %w", err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}
	return db, nil
}
