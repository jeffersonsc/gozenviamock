FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git 

WORKDIR $GOPATH/src/github.com/jeffersonsc/gozenviamock/

COPY . .

ENV TZ=America/Sao_Paulo
ENV GO111MODULE=on

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/gozenviamock github.com/jeffersonsc/gozenviamock/cmd/gozenviamock

EXPOSE 3000

ENTRYPOINT ["/go/bin/gozenviamock", "server"]