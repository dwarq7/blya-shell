GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=blya
DIR=$(BINARY_NAME)-shell


all: build

build:
	$(GOBUILD) -o $(BINARY_NAME)

test: 
	$(GOTEST) -v

clean: 
	$(GOCLEAN)
	rm -rf $(DIR)

run:
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME)

deps:
	$(GOGET) .

install:
	# main binary
	install -D -m755 "./$(BINARY_NAME)" "/usr/bin/$(BINARY_NAME)"


uninstall:
	rm -f "/usr/bin/$(BINARY_NAME)"
