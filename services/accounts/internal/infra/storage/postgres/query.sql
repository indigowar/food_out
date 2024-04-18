
-- name: GetAccountByID :one
SELECT * FROM accounts WHERE id = sqlc.arg(id)::uuid;

-- name: GetAccountByPhone :one
SELECT * FROM accounts WHERE phone = sqlc.arg(phone)::string;

-- name: GetAllAccounts :many
SELECT * FROM accounts;

-- name: InsertAccount :exec
INSERT INTO accounts (id, phone, password, name, profile_picture)
VALUES (
    sqlc.arg(id)::uuid,
    sqlc.arg(phone)::string,
    sqlc.arg(password)::string,
    sqlc.narg(name),
    sqlc.narg(profile_picture)
);

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = sqlc.arg(id)::uuid;

-- name: UpdateAccount :exec
UPDATE accounts
    SET
        phone = sqlc.arg(phone)::string,
        password = sqlc.arg(password)::string,
        name = sqlc.narg(name)::string,
        profile_picture = sqlc.narg(profile_picture)::string
    WHERE id = sqlc.arg(id)::uuid
;
