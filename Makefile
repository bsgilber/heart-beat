GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BUILDDIR=dist
BINARYNAME=main
COVERPROFILE=coverage.out

build:
	# DISABLE CGO needed to avoid dynamic linking in the network setup
	# https://stackoverflow.com/a/36308464
	CGO_ENABLED=0 $(GOBUILD) -o $(BUILDDIR)/$(BINARYNAME) cmd/api/main.go

test:
	mkdir -p $(BUILDDIR)
	$(GOTEST) -coverprofile=$(COVERPROFILE) -outputdir=$(BUILDDIR) -v ./...
	go tool cover -func $(BUILDDIR)/$(COVERPROFILE)

clean:
	rm -rf $(BUILDDIR)

local-build:
	docker build . -t health-check:test

local-test:
	make local-build
	docker-compose up
