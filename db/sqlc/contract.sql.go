// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: contract.sql

package db

import (
	"context"
)

const addContract = `-- name: AddContract :one
INSERT INTO contracts (name, network, address, is_contract, chain_id)
VALUES ($1, $2, LOWER($5), $3, $4) ON CONFLICT (address) DO
UPDATE SET name = $1, network = $2, is_contract = $3, chain_id = $4
    RETURNING id, name, is_contract, chain_id, notification, address, network, created_at
`

type AddContractParams struct {
	Name       string `json:"name"`
	Network    string `json:"network"`
	IsContract bool   `json:"is_contract"`
	ChainID    string `json:"chain_id"`
	Address    string `json:"address"`
}

func (q *Queries) AddContract(ctx context.Context, arg AddContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, addContract,
		arg.Name,
		arg.Network,
		arg.IsContract,
		arg.ChainID,
		arg.Address,
	)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsContract,
		&i.ChainID,
		&i.Notification,
		&i.Address,
		&i.Network,
		&i.CreatedAt,
	)
	return i, err
}

const deleteContract = `-- name: DeleteContract :exec
DELETE
FROM contracts
WHERE id = $1
`

func (q *Queries) DeleteContract(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteContract, id)
	return err
}

const getContract = `-- name: GetContract :one
SELECT id, name, is_contract, chain_id, notification, address, network, created_at
FROM contracts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetContract(ctx context.Context, id int64) (Contract, error) {
	row := q.db.QueryRowContext(ctx, getContract, id)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsContract,
		&i.ChainID,
		&i.Notification,
		&i.Address,
		&i.Network,
		&i.CreatedAt,
	)
	return i, err
}

const getContractByAddress = `-- name: GetContractByAddress :one
SELECT id, name, is_contract, chain_id, notification, address, network, created_at
FROM contracts
WHERE LOWER(address) = LOWER($1) LIMIT 1
`

func (q *Queries) GetContractByAddress(ctx context.Context, address string) (Contract, error) {
	row := q.db.QueryRowContext(ctx, getContractByAddress, address)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsContract,
		&i.ChainID,
		&i.Notification,
		&i.Address,
		&i.Network,
		&i.CreatedAt,
	)
	return i, err
}

const listContracts = `-- name: ListContracts :many
SELECT id, name, is_contract, chain_id, notification, address, network, created_at
FROM contracts
ORDER BY id LIMIT $1
OFFSET $2
`

type ListContractsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListContracts(ctx context.Context, arg ListContractsParams) ([]Contract, error) {
	rows, err := q.db.QueryContext(ctx, listContracts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Contract{}
	for rows.Next() {
		var i Contract
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.IsContract,
			&i.ChainID,
			&i.Notification,
			&i.Address,
			&i.Network,
			&i.CreatedAt,
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

const listTokensInContract = `-- name: ListTokensInContract :many
SELECT id, contract_id, name, symbol, token, decimals, updated_at
FROM tokens_contract
WHERE contract_id = $3 LIMIT $1
OFFSET $2
`

type ListTokensInContractParams struct {
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
	ContractID int64 `json:"contract_id"`
}

func (q *Queries) ListTokensInContract(ctx context.Context, arg ListTokensInContractParams) ([]TokensContract, error) {
	rows, err := q.db.QueryContext(ctx, listTokensInContract, arg.Limit, arg.Offset, arg.ContractID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TokensContract{}
	for rows.Next() {
		var i TokensContract
		if err := rows.Scan(
			&i.ID,
			&i.ContractID,
			&i.Name,
			&i.Symbol,
			&i.Token,
			&i.Decimals,
			&i.UpdatedAt,
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

const updateContract = `-- name: UpdateContract :one
UPDATE contracts
SET name = $2
WHERE id = $1 RETURNING id, name, is_contract, chain_id, notification, address, network, created_at
`

type UpdateContractParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateContract(ctx context.Context, arg UpdateContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, updateContract, arg.ID, arg.Name)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsContract,
		&i.ChainID,
		&i.Notification,
		&i.Address,
		&i.Network,
		&i.CreatedAt,
	)
	return i, err
}

const updateNotificationContract = `-- name: UpdateNotificationContract :one
UPDATE contracts
SET notification = $2
WHERE id = $1 RETURNING id, name, is_contract, chain_id, notification, address, network, created_at
`

type UpdateNotificationContractParams struct {
	ID           int64              `json:"id"`
	Notification NotificationStatus `json:"notification"`
}

func (q *Queries) UpdateNotificationContract(ctx context.Context, arg UpdateNotificationContractParams) (Contract, error) {
	row := q.db.QueryRowContext(ctx, updateNotificationContract, arg.ID, arg.Notification)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsContract,
		&i.ChainID,
		&i.Notification,
		&i.Address,
		&i.Network,
		&i.CreatedAt,
	)
	return i, err
}
