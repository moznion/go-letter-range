.PHONY: check test lint vet fmt-check fmt

check: test lint vet fmt-check

test:
	go test -v -cover ./...

lint:
	golint -set_exit_status ./...

vet:
	go vet ./...

fmt-check:
	gofmt -l -s letterrange/*.go | grep [^*][.]go$$; \
	EXIT_CODE=$$?; \
	if [ $$EXIT_CODE -eq 0 ]; then exit 1; fi

fmt:
	gofmt -w -s letterrange/*.go

