-- name: CreateUser :one
INSERT INTO USERS (
    FRIST_NAME,
    LAST_NAME,
    EMAIL,
    PASSWORD_HASH,
    STATUS
) VALUES(
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *;

-- name: UserLogin :one
SELECT
    PASSWORD_HASH
FROM 
    USERS
WHERE
    EMAIL = $1 LIMIT 1;

-- name: GetUserData :one
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
    UPDATED_AT = $3
WHERE
    EMAIL = $1 RETURNING *;

-- name: ResetEmail :one
UPDATE USERS
SET
    EMAIL = $2,
    UPDATED_AT = $3
WHERE
    phone_number = $1 RETURNING *;

-- name: ResetPhoneNumber :one
UPDATE USERS
SET
    phone_number = $2,
    UPDATED_AT = $3
WHERE
    EMAIL = $1 RETURNING *;;
