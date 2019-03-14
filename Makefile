install: tolog_set_active_todo.go
	mkdir -p bin
	go get -d
	go build -o bin/tolog_set_active_todo tolog_set_active_todo.go
