-- name: InsertShareKey1 :exec
INSERT INTO share_keys_1 (session_id, encrypted_share)
VALUES ($1, $2)
ON CONFLICT (session_id) DO UPDATE SET encrypted_share = EXCLUDED.encrypted_share;

-- name: InsertShareKey2 :exec
INSERT INTO share_keys_2 (session_id, encrypted_share)
VALUES ($1, $2)
ON CONFLICT (session_id) DO UPDATE SET encrypted_share = EXCLUDED.encrypted_share;

-- name: GetShareKey1 :one
SELECT encrypted_share FROM share_keys_1 WHERE session_id = $1;

-- name: GetShareKey2 :one
SELECT encrypted_share FROM share_keys_2 WHERE session_id = $1;