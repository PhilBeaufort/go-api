# Example GO API

This is a Tuto / Example for a GO API with Gin  
Reference **[Go API Tutorial - Make An API With Go](https://youtu.be/bj77B59nkTQ?si=I45DJZcSS-l00Qn4)**

Check my article for this go-api : [Golang API](https://docs.beaufort.dev/posts/golang-api/)

## Init Steps

1. Install [Go](https://go.dev/dl/)
1. Create Folder and `main.go`
1. Init  `mod` file `go mod init go-api`
1. Get dependances:

```ps1
go get github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go get github.com/jackc/pgx/v5
go get github.com/jackc/pgx/v5/pgxpool
go get github.com/lib/pq
```

## Dev Start

1. Get dependancies : `go mod download`
1. Remove unused imports : `go mod tidy`
1. Start `go run main.go`
1. Dev api will start at [http://localhost:8080](http://localhost:8080)

## Swagger doc

- Create/regenerate swagger `docs/` folder with the command `swag init`
- Sagger doc available here [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Docker Compose

Build and start app : `docker-compose up --build`

> `.env` is only used locally


## SQLC

Librairy used to generate GO code for SQL databases with SQL code of the schema and queries [sqlc.dev](https://sqlc.dev/)

- SQLC
- Postgres (pgx/v5)

Generate code : `sqlc generate`

### Install migration CLI and Go package

Dependancies :

```ps1
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/golang-migrate/migrate/v4/database/postgres
go get -u github.com/golang-migrate/migrate/v4/source/file
```

Runs all migrations in `./migrations` folder.
