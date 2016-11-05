CWD = $(shell pwd)

all: dist/epigeon

dist/epigeon:
	docker run \
	    --rm \
	    -e CGO_ENABLED=0 \
	    -v ${CWD}:/go/src/github.com/nathan-osman/epigeon \
	    -v ${CWD}/dist:/go/bin \
	    -w /go/src/github.com/nathan-osman/epigeon \
	    golang:latest \
	    go get ./...

clean:
	@rm dist/epigeon

.PHONY: clean
