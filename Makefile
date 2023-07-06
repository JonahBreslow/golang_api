BINARY_NAME=gobank

PHONY: build
build:
	@go build -o bin/$(BINARY_NAME) 

PHONY: run
run: build 
	@./bin/$(BINARY_NAME)

PHONY: test
test:
	@go test -v ./...
	
PHONY: run_db
run_db:
	@docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres  

PHONY: remove_db
remove_db: kill_db
	@docker rm some-postgres

PHONY: kill_db
kill_db:
	@docker kill some-postgres