.PHONY: build autoformat lint test clean run image install  # TODO: make this .PHONY easier to keep up to date...

build:
	go build \
		-v \
		-ldflags "-X github.com/jameshiew/kenv/cmd.Version=dev-$(shell date +%s)" \
		-o kenv main.go

autoformat:
	gofmt -w .

lint:
	docker run --rm -v $(shell pwd):/goapp -e RUN=1 -e REPO=github.com/jameshiew/kenv golangci/build-runner goenvbuild

test:
	go test -race -v ./...
	bats test

test-e2e-image: image
	docker build -f test/e2e/Dockerfile -t docker.pkg.github.com/jameshiew/kenv/kenv-test .

test-e2e: test-e2e-image
	docker run docker.pkg.github.com/jameshiew/kenv/kenv-test

clean:
	go clean

run: build
	./kenv --version

image:
	docker build -f build/package/Dockerfile -t docker.pkg.github.com/jameshiew/kenv/kenv .

install:
	go install \
		-v \
		-ldflags "-X github.com/jameshiew/kenv/cmd.Version=dev-$(shell date +%s)"
