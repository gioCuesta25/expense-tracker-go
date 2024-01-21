package db

import (
	"context"
	"database/sql"
	"fmt"
)

/*
Store proveerá todas las funciones necesarias para
ejecutar una query individualmente y combinaciones de ellas
*/
type Store struct {
	*Queries
	db *sql.DB
}

// Crea un nuevo store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Función para ejecutar transacciones en la base de datos
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		rbError := tx.Rollback()
		if rbError != nil {
			return fmt.Errorf("tx error: %v, rollback error: %v", err, rbError)
		}

		return err
	}

	return tx.Commit()
}
