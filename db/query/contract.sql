-- name: AddContract :one
INSERT INTO contracts (name, network, address, is_contract, chain_id, notification)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetContract :one
SELECT *
FROM contracts
WHERE id = $1
LIMIT 1;

-- name: GetContractByAddress :one
SELECT *
FROM contracts
WHERE LOWER(address) = LOWER(sqlc.arg('address'))
LIMIT 1;

-- name: ListContracts :many
SELECT *
FROM contracts
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: ListTokensInContract :many
SELECT *
FROM tokens_contract
WHERE contract_id = $3
LIMIT $1 OFFSET $2;

-- name: UpdateContract :one
UPDATE contracts
SET name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateNotificationContract :one
UPDATE contracts
SET notification = $2
WHERE id = $1
RETURNING *;

-- name: DeleteContract :exec
DELETE
FROM contracts
WHERE id = $1;
