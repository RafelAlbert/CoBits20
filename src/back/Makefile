

.PHONY: all run deps build 

all: clean build run

BIN_NAME=cobits20.backend

run:
	./$(BIN_NAME)


DEP_LOCATION='github.com/gorilla/mux github.com/gomodule/oauth1/oauth'

#deps: 
#	for i in $(DEP_LOCATION); do go get $i; done

build:
	go build -o $(BIN_NAME)

clean:
	- rm -f $(BIN_NAME)
