-- name: CreateUser :one
INSERT INTO USERS (
    user_id,
    FRIST_NAME,
    LAST_NAME,
    EMAIL,
    PHONE_NUMBER,
    PASSWORD_HASH,
    STATUS
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
) RETURNING *;

-- name: UserLogin :one
SELECT
    PASSWORD_HASH
FROM 
    USERS
WHERE
    EMAIL = $1 LIMIT 1;

-- name: GetUserById :one
SELECT
    frist_name,
    last_name,
    email,
    phone_number
FROM 
    USERS
WHERE
    user_id = $1;

-- name: ResetPassword :one
UPDATE USERS
SET
    PASSWORD_HASH = $2,
    UPDATED_AT = NOW()
WHERE
    EMAIL = $1 RETURNING *;

-- name: ResetEmail :one
UPDATE USERS
SET
    EMAIL = $2,
    UPDATED_AT = NOW()
WHERE
    phone_number = $1 RETURNING *;

-- name: ResetPhoneNumber :one
UPDATE USERS
SET
    phone_number = $2,
    UPDATED_AT = NOW()
WHERE
    EMAIL = $1 RETURNING *;;
