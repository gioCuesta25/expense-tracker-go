// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Category struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	IsForIncomes bool      `json:"is_for_incomes"`
	CreatedAt    time.Time `json:"created_at"`
}

type Entry struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	CategoryID  int64     `json:"category_id"`
	TypeID      int64     `json:"type_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type EntryType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
