.PHONY: build autoformat lint test clean run image install  # TODO: make this .PHONY easier to keep up to date...

build:
	go build \
		-v \
		-ldflags "-X main.version=dev-$(shell date +%s)" \
		-o kenv main.go

autoformat:
	gofmt -w .

lint-editorconfig:
	docker run --rm -v $(shell pwd):/app/code odannyc/eclint check

lint-dockerfiles:
	find . -name Dockerfile | xargs -I {} sh -c 'docker run --rm -i hadolint/hadolint:v1.17.4-0-g43bca62-debian < {}'

lint-go:
	docker run --rm -v $(shell pwd):/goapp -e RUN=1 -e REPO=github.com/jameshiew/kenv golangci/build-runner goenvbuild

lint: lint-dockerfiles lint-editorconfig lint-go

test:
	go test -race -v ./...
	bats test

test-e2e-image: image
	cd test/e2e && docker build -t docker.pkg.github.com/jameshiew/kenv/kenv-test .

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
		-ldflags "-X main.version=dev-$(shell date +%s)"

test-release:
	goreleaser --snapshot --skip-publish --rm-dist
