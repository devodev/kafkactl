FROM golang:1.17

ARG VERSION=0.0.0-development-SNAPSHOT
ARG GOOS=linux
ARG GOARCH=amd64

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
RUN go build \
		-ldflags=" \
			-X 'github.com/devodev/kafkactl/internal/kafkactl.Version=${VERSION}' \
			-X 'github.com/devodev/kafkactl/internal/kafkactl.Platform=${GOOS}-${GOARCH}'" \
        -o /usr/local/bin/kafkactl \
        ./cmd/kafkactl

RUN go build -o /usr/local/bin/gendoc ./cmd/gendoc

COPY docker/docker-entrypoint.sh /

ENTRYPOINT [ "/docker-entrypoint.sh" ]
