BINARY_NAME=gobank

build:
	@go build -o bin/$(BINARY_NAME) 

run: build 
	@./bin/$(BINARY_NAME)

test:
	@go test -v ./...
	
run_db:
	@docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres  

remove_db: kill_db
	@docker rm some-postgres

kill_db:
	@docker kill some-postgres