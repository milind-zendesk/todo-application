
.DEFAULT_GOAL := run

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

get-todos:
	curl localhost:8080/todos | jq .

add-todo:
	curl -X POST localhost:8080/insert_todo -d '{"title":"Sleep","status":"Pending"}'

vendor:
	go mod vendor

test: vendor
	go test -cover -mod=vendor ./...