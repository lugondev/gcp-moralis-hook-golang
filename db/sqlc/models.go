// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type NotificationStatus string

const (
	NotificationStatusDisable NotificationStatus = "disable"
	NotificationStatusEnable  NotificationStatus = "enable"
)

func (e *NotificationStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = NotificationStatus(s)
	case string:
		*e = NotificationStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for NotificationStatus: %T", src)
	}
	return nil
}

type NullNotificationStatus struct {
	NotificationStatus NotificationStatus
	Valid              bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullNotificationStatus) Scan(value interface{}) error {
	if value == nil {
		ns.NotificationStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.NotificationStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullNotificationStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.NotificationStatus, nil
}

type Contract struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	IsContract   bool               `json:"is_contract"`
	ChainID      string             `json:"chain_id"`
	Notification NotificationStatus `json:"notification"`
	Address      string             `json:"address"`
	Network      string             `json:"network"`
	CreatedAt    time.Time          `json:"created_at"`
}

type Email struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type EmailsContract struct {
	ID         int64     `json:"id"`
	EmailID    int64     `json:"email_id"`
	ContractID int64     `json:"contract_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type TokensContract struct {
	ID         int64     `json:"id"`
	ContractID int64     `json:"contract_id"`
	Name       string    `json:"name"`
	Symbol     string    `json:"symbol"`
	Address    string    `json:"address"`
	Decimals   int64     `json:"decimals"`
	UpdatedAt  time.Time `json:"updated_at"`
}
