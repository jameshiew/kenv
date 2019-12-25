FROM golang:1.13.5-buster

WORKDIR /go/src/github.com/jameshiew/kenv
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
COPY cmd cmd
COPY internal internal
COPY main.go main.go

RUN go install -v ./...

ENTRYPOINT ["kenv"]