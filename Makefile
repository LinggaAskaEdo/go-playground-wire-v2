.PHONY: build
build:
	@wire ./src/cmd && \
	go mod tidy && \
	go build -o ./build/app ./src/cmd

.PHONY: vendor
vendor:
	@go mod vendor

.PHONY: run
run:
	@./build/app