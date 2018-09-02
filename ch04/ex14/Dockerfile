FROM golang:alpine
MAINTAINER "patrick.bucher@stud.hslu.ch"

ENV SOURCE_DIR "/go/src/gopl.io/ch04/ex14"
RUN mkdir -p ${SOURCE_DIR}
COPY cmd "${SOURCE_DIR}/cmd"
COPY *.go "${SOURCE_DIR}/"
WORKDIR ${SOURCE_DIR}

RUN go build -o issueserver cmd/issueserver.go
RUN cp issueserver /go/bin/

EXPOSE 8000
