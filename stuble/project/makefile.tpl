run:
	go run -tags dev .

build:
	go build -o bin/{{ toKebab .ProjectName }} .

swagger-gen:
	swag init -g cmd.go \
		--dir src/module/{{ toSnake .ModuleName }} \
		--output src/docs/{{ toSnake .ModuleName }} \
		--parseDependency \
		--parseInternal

.PHONY: run build swagger-gen
