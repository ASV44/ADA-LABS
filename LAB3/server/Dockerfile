FROM golang:alpine AS builder

# Set environment variables

ENV APP ADA-LAB3
ENV WORKDIR ${GOPATH}/src/github.com/ASV44/${APP}

#Change container work directory

WORKDIR $WORKDIR

# Install git.
# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git && \
    go get -u github.com/golang/dep/cmd/dep

#Copy dependecies list files

COPY Gopkg.toml $WORKDIR/
COPY Gopkg.lock $WORKDIR/

#Copy source code to build image

COPY db $WORKDIR/db
COPY server $WORKDIR/server

# Fetch dependencies.
# Using dep.

RUN set -x && \
    dep ensure -v

# Build the binary.

RUN go build -o /go/bin/main server/main.go

# Run the binary.
ENTRYPOINT ["/go/bin/main"]