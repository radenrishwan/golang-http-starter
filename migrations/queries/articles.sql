-- name: GetArticleBySlug :one
SELECT * FROM articles WHERE slug = $1 LIMIT 1;

-- name: GetArticleWithOffset :many
SELECT * FROM articles ORDER BY created_at DESC OFFSET $1 LIMIT $2;

-- name: CreateArticle :one
INSERT INTO articles (title, slug, body) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateArticle :one
UPDATE articles SET title = $1, slug = $2, body = $3 WHERE id = $4 RETURNING *;

-- name: DeleteArticle :one
DELETE FROM articles WHERE id = $1 RETURNING *;
