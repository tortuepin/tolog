bins = tolog_set_active_todo tolog_tag_collect


all: $(bins)

$(bins): 
	mkdir -p bin
	go build -o bin/$@ $@.go
