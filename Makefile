include Makefile.tool

.DEFAULT_GOAL := run

MOCK_FILES = $(shell grep -l -R "go:generate mockgen" endpoint/api/.)

##@ Bootstrap the project
.PHONY: bootstrap

bootstrap: ## Run the bootstrap script and check if required tools are available, installing them if possible
	./script/local-dev/bootstrap.sh

##@ Run the project locally
.PHONY: run start-mysql

run: start-mysql
	go run main.go

start-mysql: 
	./script/local-dev/start-mysql.sh

generate-fake-data: ##To fill the user and todo tables with fake/dummy data
	curl -X POST localhost:8080/bulk_load

get-todos: ##To get all records from todo
	curl localhost:8080/todos | jq .

get-users: ##To get all records from user
	curl localhost:8080/users | jq .

add-todo: ##To add new entry in todo table
	curl -X POST localhost:8080/insert_todo -d '{"title":"Sleep","status":"Pending"}'

vendor:
	go mod vendor

test: vendor ##To run the tests
	go test -cover -mod=vendor ./...

generate-mocks: $(MOCKGEN) $(MOCK_FILES) ## Generate test mock files with mockgen
	@PATH=$(shell pwd)/$(TOOLS_BIN):$$PATH go generate ./...
	touch generate-mocks
