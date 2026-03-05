install:
	go install github.com/Mirsadikovv/gengo@latest
	export PATH=$$PATH:$$(go env GOPATH)/bin

install-local:
	go install .
	export PATH=$$PATH:$$(go env GOPATH)/bin

.PHONY: install install-local
