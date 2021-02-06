GO=go
files_clean=get get.zip insert insert.zip

test:
	$(GO) test ./...

compile: test
	$(GO) build pkg/get-function/get.go
	$(GO) build pkg/insert-function/insert.go

install: compile
	zip get.zip get
	zip insert.zip insert

clean:
	rm $(files_clean)

