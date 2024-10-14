
-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = sqlc.arg(id)::uuid;

-- name: GetAccountByPhone :one
SELECT * FROM accounts WHERE phone = sqlc.arg(phone)::text;

-- name: GetAllAccounts :many
SELECT * FROM accounts;

-- name: InsertAccount :exec
INSERT INTO accounts (id, phone, password, name, profile)
VALUES (
    sqlc.arg(id)::uuid,
    sqlc.arg(phone)::text,
    sqlc.arg(password)::text,
    sqlc.narg(name),
    sqlc.narg(profile)
);

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = sqlc.arg(id)::uuid;

-- name: UpdateAccount :exec
UPDATE accounts
    SET
        phone = sqlc.arg(phone)::text,
        password = sqlc.arg(password)::text,
        name = sqlc.narg(name)::text,
        profile = sqlc.narg(profile)::text
    WHERE id = sqlc.arg(id)::uuid
;
