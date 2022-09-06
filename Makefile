GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

BUILDDIR=dist
BINARYNAME=main
COVERPROFILE=coverage.out

build:
	$(GOBUILD) -o $(BUILDDIR)/$(BINARYNAME) cmd/api/main.go

test:
	mkdir -p $(BUILDDIR)
	$(GOTEST) -coverprofile=$(COVERPROFILE) -outputdir=$(BUILDDIR) -v ./...
	go tool cover -func $(BUILDDIR)/$(COVERPROFILE)

clean:
	rm -rf $(BUILDDIR)

local-run:
	docker build . -t health-check:test
	docker run -p 8080:8080 health-check:test
