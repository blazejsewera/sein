ARG GO_VERSION=1.24
ARG DEBIAN_VERSION=bookworm
ARG BUSYBOX_VERSION=stable-glibc
# build
FROM docker.io/library/golang:${GO_VERSION}-${DEBIAN_VERSION} AS build-env

WORKDIR /

# preload dependencies (currently there are no dependencies)
#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

COPY . .

RUN make build

# run
FROM docker.io/library/busybox:${BUSYBOX_VERSION}

COPY --from=build-env /sein /

EXPOSE 8080
STOPSIGNAL SIGINT

ENTRYPOINT ["/sein"]
