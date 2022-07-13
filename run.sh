# Start postgres locally. Create if doesn't exist.
docker run --name concordanceDB -p 5433:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=concordance -d postgres || docker start concordanceDB
# Start the server
cd ./controller
go mod tidy
go run cmd/main.go