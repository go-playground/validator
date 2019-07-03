GOCMD=go

linters-install:
	$(GOCMD) get -u github.com/alecthomas/gometalinter
	gometalinter --install

lint: linters-install
	gometalinter --vendor --disable-all --enable=vet --enable=vetshadow --enable=golint --enable=maligned --enable=megacheck --enable=ineffassign --enable=misspell --enable=errcheck --enable=goconst ./...

test:
	$(GOCMD) test -cover -race ./...

bench:
	$(GOCMD) test -bench=. -benchmem ./...

.PHONY: test lint linters-install