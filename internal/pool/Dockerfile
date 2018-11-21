# To regenerate CA bundle into pool.go
# docker build -t cabin-gen .
# docker run --rm -ti -v /go/src/github.com/nogoegst/cabin/internal/pool:/workdir cabin-gen
FROM golang:1.11-alpine3.8
RUN apk add --no-cache ca-certificates
WORKDIR /workdir
CMD ["go", "generate"]
