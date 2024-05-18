-- name: CreateClient :exec
INSERT INTO clients (
    id,
    name,
    age,
    social_info_instagram,
    social_info_linkedin,
    social_info_facebook
)
VALUES ($1, $2, $3, $4,$5,$6);

-- name: GetClients :many
SELECT * FROM clients;

-- name: GetClient :one
SELECT * FROM clients WHERE id = $1;

-- name: UpdateClient :exec
UPDATE clients
SET
    name = $1,
    age = $2,
    social_info_instagram = $3,
    social_info_linkedin = $4,
    social_info_facebook = $5
WHERE id = $6;

-- name: DeleteClient :exec
DELETE FROM clients WHERE id = $1;