.PHONY: run start-mysql

run: start-mysql
	go run main.go

start-mysql: 
	./script/local-dev/start-mysql.sh
