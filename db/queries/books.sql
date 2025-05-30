-- name: ListBooks :many
SELECT * FROM books;

-- name: GetBook :one
SELECT * FROM books WHERE id = $1;

-- name: CreateBook :one
INSERT INTO books (title, author, quantity) VALUES ($1, $2, $3)
RETURNING *;

-- name: CheckoutBook :one
UPDATE books SET quantity = quantity - 1 WHERE id = $1 AND quantity > 0 RETURNING *;

-- name: ReturnBook :one
UPDATE books SET quantity = quantity + 1 WHERE id = $1 RETURNING *;
