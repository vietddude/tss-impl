// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: share_keys.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getShareKey1 = `-- name: GetShareKey1 :one
SELECT encrypted_share FROM share_keys_1 WHERE session_id = $1
`

func (q *Queries) GetShareKey1(ctx context.Context, sessionID pgtype.UUID) ([]byte, error) {
	row := q.db.QueryRow(ctx, getShareKey1, sessionID)
	var encrypted_share []byte
	err := row.Scan(&encrypted_share)
	return encrypted_share, err
}

const getShareKey2 = `-- name: GetShareKey2 :one
SELECT encrypted_share FROM share_keys_2 WHERE session_id = $1
`

func (q *Queries) GetShareKey2(ctx context.Context, sessionID pgtype.UUID) ([]byte, error) {
	row := q.db.QueryRow(ctx, getShareKey2, sessionID)
	var encrypted_share []byte
	err := row.Scan(&encrypted_share)
	return encrypted_share, err
}

const insertShareKey1 = `-- name: InsertShareKey1 :exec
INSERT INTO share_keys_1 (session_id, encrypted_share)
VALUES ($1, $2)
ON CONFLICT (session_id) DO UPDATE SET encrypted_share = EXCLUDED.encrypted_share
`

type InsertShareKey1Params struct {
	SessionID      pgtype.UUID `json:"session_id"`
	EncryptedShare []byte      `json:"encrypted_share"`
}

func (q *Queries) InsertShareKey1(ctx context.Context, arg InsertShareKey1Params) error {
	_, err := q.db.Exec(ctx, insertShareKey1, arg.SessionID, arg.EncryptedShare)
	return err
}

const insertShareKey2 = `-- name: InsertShareKey2 :exec
INSERT INTO share_keys_2 (session_id, encrypted_share)
VALUES ($1, $2)
ON CONFLICT (session_id) DO UPDATE SET encrypted_share = EXCLUDED.encrypted_share
`

type InsertShareKey2Params struct {
	SessionID      pgtype.UUID `json:"session_id"`
	EncryptedShare []byte      `json:"encrypted_share"`
}

func (q *Queries) InsertShareKey2(ctx context.Context, arg InsertShareKey2Params) error {
	_, err := q.db.Exec(ctx, insertShareKey2, arg.SessionID, arg.EncryptedShare)
	return err
}
