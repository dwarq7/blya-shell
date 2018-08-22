GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=blya
DIR=$(BINARY_NAME)-shell

DESTDIR?=./result

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
	install -D -m755 "./$(BINARY_NAME)" "$(DESTDIR)/usr/bin/$(BINARY_NAME)"


uninstall:
	rm -f "$(DESTDIR)/usr/bin/$(BINARY_NAME)"
