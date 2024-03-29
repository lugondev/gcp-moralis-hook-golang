// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	AddChain(ctx context.Context, arg AddChainParams) (Chain, error)
	AddContract(ctx context.Context, arg AddContractParams) (Contract, error)
	AddEmail(ctx context.Context, arg AddEmailParams) (Email, error)
	AddPrivateKey(ctx context.Context, arg AddPrivateKeyParams) (PrivateKey, error)
	AddTokenToContract(ctx context.Context, arg AddTokenToContractParams) (TokensContract, error)
	AddTokenToContractByAddress(ctx context.Context, arg AddTokenToContractByAddressParams) (TokensContract, error)
	DeleteContract(ctx context.Context, id int64) error
	DeleteEmail(ctx context.Context, id int64) error
	GetChain(ctx context.Context, chainID int64) (Chain, error)
	GetContract(ctx context.Context, id int64) (Contract, error)
	GetContractByAddress(ctx context.Context, address string) (Contract, error)
	GetEmail(ctx context.Context, id int64) (Email, error)
	GetEmailForContract(ctx context.Context, contractID int64) ([]EmailsContract, error)
	GetPrivateKey(ctx context.Context, address string) (PrivateKey, error)
	ListAddresses(ctx context.Context) ([]string, error)
	ListChains(ctx context.Context) ([]Chain, error)
	ListContracts(ctx context.Context, arg ListContractsParams) ([]Contract, error)
	ListEmails(ctx context.Context, arg ListEmailsParams) ([]Email, error)
	ListEmailsSubscription(ctx context.Context, contractid int64) ([]Email, error)
	ListEmailsSubscriptionByAddress(ctx context.Context, contractaddress string) ([]Email, error)
	ListTokenInContract(ctx context.Context, contractID int64) ([]TokensContract, error)
	ListTokenInContractByAddress(ctx context.Context, address string) ([]TokensContract, error)
	ListTokensInContract(ctx context.Context, arg ListTokensInContractParams) ([]TokensContract, error)
	MapEmailContract(ctx context.Context, arg MapEmailContractParams) (MapEmailContractRow, error)
	RemoveTokenFromContract(ctx context.Context, arg RemoveTokenFromContractParams) error
	RemoveTokenFromContractByAddress(ctx context.Context, arg RemoveTokenFromContractByAddressParams) error
	UpdateContract(ctx context.Context, arg UpdateContractParams) (Contract, error)
	UpdateEmail(ctx context.Context, arg UpdateEmailParams) (Email, error)
	UpdateNotificationContract(ctx context.Context, arg UpdateNotificationContractParams) (Contract, error)
}

var _ Querier = (*Queries)(nil)
