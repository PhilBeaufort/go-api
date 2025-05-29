# Example GO API

This is a Tuto / Example for a GO API with Gin  
Reference **[Go API Tutorial - Make An API With Go](https://youtu.be/bj77B59nkTQ?si=I45DJZcSS-l00Qn4)**

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
```

## Dev Start

1. Start `go run main.go`
1. Dev api will start at [http://localhost:8080](http://localhost:8080)

## Swagger doc

- Create/regenerate swagger `docs/` folder with the command `swag init`
- Sagger doc available here [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Docker container

1. Build image : `docker build -t go-api .`
1. Run image : `docker run --rm --name go-api -p 8080:8080 go-api`
