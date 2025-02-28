// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: addresses.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createAddress = `-- name: CreateAddress :one
INSERT INTO ADDRESSES (
    ADDRESS_ID,
    USER_ID,
    ADDRESS,
    CITY,
    STATE,
    POSTAL_CODE,
    COUNTRY,
    IS_DEFAULT
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING address_id, user_id, address, city, state, postal_code, country, is_default, created_at, updated_at
`

type CreateAddressParams struct {
	AddressID  uuid.UUID `json:"address_id"`
	UserID     uuid.UUID `json:"user_id"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	IsDefault  bool      `json:"is_default"`
}

func (q *Queries) CreateAddress(ctx context.Context, arg CreateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, createAddress,
		arg.AddressID,
		arg.UserID,
		arg.Address,
		arg.City,
		arg.State,
		arg.PostalCode,
		arg.Country,
		arg.IsDefault,
	)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteAddress = `-- name: DeleteAddress :exec
DELETE FROM ADDRESSES
WHERE
    USER_ID = $1 AND ADDRESS_ID = $2
`

type DeleteAddressParams struct {
	UserID    uuid.UUID `json:"user_id"`
	AddressID uuid.UUID `json:"address_id"`
}

func (q *Queries) DeleteAddress(ctx context.Context, arg DeleteAddressParams) error {
	_, err := q.db.Exec(ctx, deleteAddress, arg.UserID, arg.AddressID)
	return err
}

const getAddressesByUserID = `-- name: GetAddressesByUserID :many
SELECT 
    ADDRESS_ID,
    ADDRESS,
    CITY,
    STATE,
    POSTAL_CODE,
    COUNTRY,
    IS_DEFAULT
FROM 
    ADDRESSES 
WHERE 
    USER_ID = $1
`

type GetAddressesByUserIDRow struct {
	AddressID  uuid.UUID `json:"address_id"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	IsDefault  bool      `json:"is_default"`
}

func (q *Queries) GetAddressesByUserID(ctx context.Context, userID uuid.UUID) ([]GetAddressesByUserIDRow, error) {
	rows, err := q.db.Query(ctx, getAddressesByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAddressesByUserIDRow{}
	for rows.Next() {
		var i GetAddressesByUserIDRow
		if err := rows.Scan(
			&i.AddressID,
			&i.Address,
			&i.City,
			&i.State,
			&i.PostalCode,
			&i.Country,
			&i.IsDefault,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetDefaultAddress = `-- name: ResetDefaultAddress :one
UPDATE ADDRESSES
SET 
    IS_DEFAULT = FALSE
WHERE 
    ADDRESS_ID = $1 RETURNING address_id, user_id, address, city, state, postal_code, country, is_default, created_at, updated_at
`

func (q *Queries) ResetDefaultAddress(ctx context.Context, addressID uuid.UUID) (Address, error) {
	row := q.db.QueryRow(ctx, resetDefaultAddress, addressID)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const setDefaultAddress = `-- name: SetDefaultAddress :one
UPDATE ADDRESSES
SET
    IS_DEFAULT = TRUE
WHERE
    ADDRESS_ID = $1 RETURNING address_id, user_id, address, city, state, postal_code, country, is_default, created_at, updated_at
`

func (q *Queries) SetDefaultAddress(ctx context.Context, addressID uuid.UUID) (Address, error) {
	row := q.db.QueryRow(ctx, setDefaultAddress, addressID)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateAddress = `-- name: UpdateAddress :one
UPDATE ADDRESSES
SET
    ADDRESS = $2,
    CITY = $3,
    STATE = $4,
    POSTAL_CODE = $5,
    COUNTRY = $6,
    IS_DEFAULT = $7,
    UPDATED_AT = NOW()
WHERE
    ADDRESS_ID = $1 RETURNING address_id, user_id, address, city, state, postal_code, country, is_default, created_at, updated_at
`

type UpdateAddressParams struct {
	AddressID  uuid.UUID `json:"address_id"`
	Address    string    `json:"address"`
	City       string    `json:"city"`
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	IsDefault  bool      `json:"is_default"`
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) (Address, error) {
	row := q.db.QueryRow(ctx, updateAddress,
		arg.AddressID,
		arg.Address,
		arg.City,
		arg.State,
		arg.PostalCode,
		arg.Country,
		arg.IsDefault,
	)
	var i Address
	err := row.Scan(
		&i.AddressID,
		&i.UserID,
		&i.Address,
		&i.City,
		&i.State,
		&i.PostalCode,
		&i.Country,
		&i.IsDefault,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
