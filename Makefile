GO = go
BIN = patates

PREFIX = /usr/local

.PHONY : all build clean install uninstall

all: build

build:
	$(GO) build -o $(BIN)

install: build
	mkdir -p $(PREFIX)/bin
	cp $(BIN) $(PREFIX)/bin

uninstall:
	rm $(PREFIX)/bin/$(BIN)

clean:
	rm $(BIN)
