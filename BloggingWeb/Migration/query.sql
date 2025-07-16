-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetAllBlogs :many
SELECT * FROM blogs;

-- name: PostBlog :one
INSERT INTO blogs (title, content, user_id, first_name, last_name, full_name,email)
VALUES ($1, $2, $3 , $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1;

-- name: InsertUser :one
INSERT INTO users (full_name, first_name, last_name, email, role, status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: UpdateUserBlogCount :one
UPDATE users
    SET blogs_uploaded = blogs_uploaded + $1
    WHERE user_id = $2
RETURNING *;

-- name: GetUserByUid :one
SELECT * FROM users
    WHERE user_id = $1;

-- name: CheckUserExists :one
SELECT EXISTS (
  SELECT * FROM users WHERE email = $1
);

-- name: GetBlogByID :one
SELECT * FROM blogs
    WHERE blog_id = $1;

-- name: GetUserBlogs :many
SELECT * FROM blogs
    WHERE user_id = $1;

-- name: UpdateBlogsContent :one
UPDATE blogs
    SET content = $1, title = $2
    WHERE blog_id = $3
RETURNING *;

-- name: DeleteBlogByID :exec
DELETE FROM blogs
    WHERE blog_id = $1;

-- name: UpdateUserDetails :one
UPDATE users
    SET full_name = $1, first_name = $2, last_name = $3, date_of_birth = $4, user_address = $5, updated_at = $6
    WHERE user_id = $7
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users 
    WHERE email = $1;

-- name: GetUserCredByEmail :one
SELECT * FROM user_credentials
    WHERE email = $1;

-- name: UpdateSessionTokenAndAgent :one
UPDATE user_sessions
    SET token = $1, user_agent = $2, created_at = $3, status = $4
    WHERE email = $5
RETURNING *;

-- name: CheckUserRegistration :one
SELECT EXISTS (
  SELECT * FROM user_credentials
    WHERE email = $1
);

-- name: SaveUserCredentials :exec
INSERT INTO 
    user_credentials (uid,email, password, role, first_name, last_name) 
    VALUES ($1, $2, $3, $4, $5, $6);

-- name: SaveUserSession :exec
INSERT INTO user_sessions (uid, email, token, user_agent, role)
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateUserSessionStatus :exec
UPDATE user_sessions
    SET status = $1
    WHERE uid = $2;

-- name: UpdateUserStatus :exec
UPDATE users
    SET status = $1
    WHERE user_id = $2;