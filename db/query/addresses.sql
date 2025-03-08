-- name: CreateAddress :one
INSERT INTO ADDRESSES (
    ADDRESS_ID,
    USER_ID,
    flat_floor,
    building,
    street,
    district,
    region,
    country,
    zip_code,
    IS_DEFAULT
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
) RETURNING *;

-- name: GetAddressesByUserID :many
SELECT 
    *
FROM 
    ADDRESSES 
WHERE 
    USER_ID = $1;

-- name: UpdateAddress :exec
UPDATE ADDRESSES
SET
    flat_floor = $3,
    building = $4,
    street = $5,
    district = $6,
    region = $7,
    country = $8,
    zip_code = $9,
    IS_DEFAULT = $10,
    UPDATED_AT = NOW()
WHERE
    ADDRESS_ID = $1 AND user_id = $2;

-- name: SetAllAddresstoFalse :exec
UPDATE ADDRESSES
SET 
    IS_DEFAULT = FALSE
WHERE 
    user_id = $1;

-- name: SetDefaultAddress :exec
UPDATE ADDRESSES
SET
    IS_DEFAULT = TRUE
WHERE
    ADDRESS_ID = $1 AND user_id = $2;

-- name: DeleteAddress :exec
DELETE FROM ADDRESSES
WHERE
    ADDRESS_ID = $1 AND USER_ID = $2;
