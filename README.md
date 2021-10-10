# GoCardAuctionSample
practice go

# Require
## go
```shell
go version                        
# go version go1.15.5 darwin/amd64
```

## Docker Desktop
Ref: https://www.docker.com/products/docker-desktop
version: 3.6.0（3.6.0.5487）

## PostgreSQL
```shell
docker run -p 5432:5432 --name eth-block-indexer -e POSTGRES_PASSWORD=mysecretpassword -d postgres
```
1. user name: postgres
1. password: mysecretpassword
1. db name: postgres

**PS: need check port 5432 is available**

# Run application
## build
```shell
# build at project root folder
go build

# run
go run main.go
```
Note:
1. Listening and serving HTTP on localhost:8080

# Call API
## API service
```shell
```

