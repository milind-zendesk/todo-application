.PHONY: run start-mysql

run: start-mysql
	go run main.go

start-mysql: 
	./script/local-dev/start-mysql.sh

get-todos:
	curl localhost:8080/todos | jq .