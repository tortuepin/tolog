bins = tolog_set_active_todo tolog_tag_collect tolog_log_search_bytag tolog_get_random


all: $(bins)

$(bins): 
	mkdir -p bin
	go build -o bin/$@ $@.go

clean:
	rm -fr bin/*
