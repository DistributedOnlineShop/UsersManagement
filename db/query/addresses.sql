-- name: CreateAddress :one
INSERT INTO ADDRESSES (
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
    $7
) RETURNING *;

-- name: GetAddresses :many
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
    USER_ID = $1;

-- name: UpdateAddress :one
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
    ADDRESS_ID = $1 RETURNING *;

-- name: ResetDefaultAddress :one
UPDATE ADDRESSES
SET 
    IS_DEFAULT = FALSE
WHERE 
    USER_ID = $1 RETURNING *;

-- name: SetDefaultAddress :one
UPDATE ADDRESSES
SET
    IS_DEFAULT = TRUE
WHERE
    USER_ID = $1 AND ADDRESS_ID = $2
    RETURNING *;

-- name: DeleteAddress :exec
DELETE FROM ADDRESSES
WHERE
    USER_ID = $1 AND ADDRESS_ID = $2;
