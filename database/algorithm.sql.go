// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: algorithm.sql

package database

import (
	"context"
)

const createAlgorithm = `-- name: CreateAlgorithm :exec
INSERT INTO algorithms (text_id, name, type, explanation, created_at, updated_at)
VALUES ($1, $2, $3, $4, timezone('utc', NOW()), timezone('utc', NOW()))
`

type CreateAlgorithmParams struct {
	TextID      string
	Name        string
	Type        string
	Explanation string
}

func (q *Queries) CreateAlgorithm(ctx context.Context, arg CreateAlgorithmParams) error {
	_, err := q.db.ExecContext(ctx, createAlgorithm,
		arg.TextID,
		arg.Name,
		arg.Type,
		arg.Explanation,
	)
	return err
}

const getAlgorithmById = `-- name: GetAlgorithmById :one
SELECT id, text_id, name, type, explanation, created_at, updated_at FROM algorithms
WHERE id = $1
`

func (q *Queries) GetAlgorithmById(ctx context.Context, id int32) (Algorithm, error) {
	row := q.db.QueryRowContext(ctx, getAlgorithmById, id)
	var i Algorithm
	err := row.Scan(
		&i.ID,
		&i.TextID,
		&i.Name,
		&i.Type,
		&i.Explanation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAlgorithmByTextId = `-- name: GetAlgorithmByTextId :one
SELECT id, text_id, name, type, explanation, created_at, updated_at FROM algorithms
WHERE text_id = $1
`

func (q *Queries) GetAlgorithmByTextId(ctx context.Context, textID string) (Algorithm, error) {
	row := q.db.QueryRowContext(ctx, getAlgorithmByTextId, textID)
	var i Algorithm
	err := row.Scan(
		&i.ID,
		&i.TextID,
		&i.Name,
		&i.Type,
		&i.Explanation,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAlgorithm = `-- name: UpdateAlgorithm :exec
UPDATE algorithms
SET text_id = $2, 
    name = $3,
    updated_at = timezone('utc', NOW()) 
WHERE id = $1
`

type UpdateAlgorithmParams struct {
	ID     int32
	TextID string
	Name   string
}

func (q *Queries) UpdateAlgorithm(ctx context.Context, arg UpdateAlgorithmParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithm, arg.ID, arg.TextID, arg.Name)
	return err
}

const updateAlgorithmExplanation = `-- name: UpdateAlgorithmExplanation :exec
UPDATE algorithms
SET explanation = $2,
    updated_at = timezone('utc', NOW()) 
WHERE id = $1
`

type UpdateAlgorithmExplanationParams struct {
	ID          int32
	Explanation string
}

func (q *Queries) UpdateAlgorithmExplanation(ctx context.Context, arg UpdateAlgorithmExplanationParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithmExplanation, arg.ID, arg.Explanation)
	return err
}

const updateAlgorithmName = `-- name: UpdateAlgorithmName :exec
UPDATE algorithms
SET name = $2, 
    updated_at = timezone('utc', NOW()) 
WHERE id = $1
`

type UpdateAlgorithmNameParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateAlgorithmName(ctx context.Context, arg UpdateAlgorithmNameParams) error {
	_, err := q.db.ExecContext(ctx, updateAlgorithmName, arg.ID, arg.Name)
	return err
}
