GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run
BINARY_FOLDER=build
SCRIPT_ID=0

all: test build
build:
	$(GOCLEAN)
	$(GOBUILD) -a -o $(BINARY_FOLDER)/nrutil .

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FOLDER)/*

run:
	$(GORUN) . getscript --id $(SCRIPT_ID)
