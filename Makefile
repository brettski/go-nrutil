GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_FOLDER=build
SCRIPT_ID=0

all: test build
build:
	$(GOBUILD) -o $(BINARY_FOLDER)/nrutil ./command/.

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FOLDER)/*

run:
	$(GORUN) ./command/. getscript --id $(SCRIPT_ID)
