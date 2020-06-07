FROM golang:alpine

RUN mkdir -p /go/src/gitlab.com/snpranav/go-api-boilerplate

WORKDIR /go/src/gitlab.com/snpranav/go-api-boilerplate

COPY api.go .

# COPY test script
RUN mkdir tests/
COPY tests/ tests/
RUN go build tests/api-check.go; rm tests/api-check.go

RUN go build api.go; rm api.go

EXPOSE 3000

ENTRYPOINT ["./api"]