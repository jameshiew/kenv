.PHONY: build autoformat lint clean run image install  # TODO: make this .PHONY easier to keep up to date...

build: lint
	go build \
		-v \
		-ldflags "-X github.com/jameshiew/kenv/cmd.Version=dev-$(shell date +%s)" \
		-o kenv main.go

autoformat:
	gofmt -w .

lint:
	gofmt -l . # TODO make this error out
	golint ./...
	go vet ./...

test-e2e-image: image
	docker build -f test/e2e/Dockerfile -t kenv-test .

test-e2e: test-e2e-image
	docker run kenv-test

clean:
	go clean

run: build
	./kenv --help

image:
	docker build -f build/package/Dockerfile -t kenv .

install: lint
	go install \
		-v \
		-ldflags "-X github.com/jameshiew/kenv/cmd.Version=dev-$(shell date +%s)"
