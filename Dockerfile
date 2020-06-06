FROM golang:alpine

RUN mkdir -p /go/src/github.com/go-api

WORKDIR /go/src/github.com/go-api

COPY api.go .

RUN go build api.go

EXPOSE 3000

ENTRYPOINT ["./api"]