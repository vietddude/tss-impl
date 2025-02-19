-- +goose Up
CREATE TABLE share_keys_1 (
    session_id UUID PRIMARY KEY,
    encrypted_share BYTEA NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE share_keys_2 (
    session_id UUID PRIMARY KEY,
    encrypted_share BYTEA NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- +goose Down
DROP TABLE share_keys_1;
DROP TABLE share_keys_2;
