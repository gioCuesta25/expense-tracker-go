// Code generated by sqlc. DO NOT EDIT.
// source: entries.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
    description, amount, category_id, type_id
) VALUES (
    $1, $2, $3, $4
)
RETURNING id, description, amount, category_id, type_id, created_at, user_id
`

type CreateEntryParams struct {
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	CategoryID  uuid.UUID `json:"category_id"`
	TypeID      uuid.UUID `json:"type_id"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry,
		arg.Description,
		arg.Amount,
		arg.CategoryID,
		arg.TypeID,
	)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Amount,
		&i.CategoryID,
		&i.TypeID,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, description, amount, category_id, type_id, created_at, user_id FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id uuid.UUID) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Description,
		&i.Amount,
		&i.CategoryID,
		&i.TypeID,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, description, amount, category_id, type_id, created_at, user_id FROM entries
ORDER BY created_at
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Entry
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Description,
			&i.Amount,
			&i.CategoryID,
			&i.TypeID,
			&i.CreatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateEntry = `-- name: UpdateEntry :exec
UPDATE entries
    set description = $2,
    amount = $3,
    category_id = $4,
    type_id = $5
WHERE id = $1
`

type UpdateEntryParams struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Amount      int64     `json:"amount"`
	CategoryID  uuid.UUID `json:"category_id"`
	TypeID      uuid.UUID `json:"type_id"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) error {
	_, err := q.db.ExecContext(ctx, updateEntry,
		arg.ID,
		arg.Description,
		arg.Amount,
		arg.CategoryID,
		arg.TypeID,
	)
	return err
}
