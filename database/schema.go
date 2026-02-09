package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func initSchema(db *sqlx.DB) error {
	tx := db.MustBeginTx(context.Background(), &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
	defer tx.Rollback()

	tables := []string{
		usersTable,
	}
	if err := executeQueryList(tx, "Tables", tables); err != nil {
		return err
	}

	indexes := []string{
		usersIndex,
	}
	if err := executeQueryList(tx, "Index", indexes); err != nil {
		return err
	}

	constraints := []string{}
	if err := executeQueryList(tx, "Constraints", constraints); err != nil {
		return err
	}

	return tx.Commit()
}

func executeQueryList(tx *sqlx.Tx, queryName string, queries []string) error {
	for i, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			return fmt.Errorf("Не удалось выполнить транзакцию %s-%d: %w", queryName, i, err)
		}
	}

	return nil
}

// SQL schemas
var usersTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(64) UNIQUE NOT NULL,
	password VARCHAR(128) NOT NULL
);
`

var usersIndex = `
CREATE INDEX IF NOT EXISTS user_username ON users (username);
`
