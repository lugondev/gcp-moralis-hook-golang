-- name: AddToken :one
INSERT INTO tokens_contract (contract_id, name, symbol, address, decimals)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: DeleteToken :exec
DELETE
FROM tokens_contract
WHERE id = $1;
