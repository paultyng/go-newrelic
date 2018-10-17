all: fix vet lint test

fix:
	go fix ./...

vet:
	go vet ./...

test:
	go test -v -cover ./...

lint:
	printf "%s\n" "$(NOVENDOR)" | xargs -I {} sh -c 'golint -set_exit_status {}'
