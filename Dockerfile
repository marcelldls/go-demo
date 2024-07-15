ARG GOLANG_VER=1.22
ARG ALPINE_VER=3.20

FROM golang:${GOLANG_VER}-alpine${ALPINE_VER} as developer

RUN go install -v github.com/go-delve/delve/cmd/dlv@latest