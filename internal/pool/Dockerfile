# To regenerate CA bundle into pool.go
# docker build -t cabin-gen .
# docker run --rm cabin-gen > pool.go
FROM golang:alpine
RUN apk add ca-certificates
WORKDIR /workdir
COPY generate.go .
CMD go generate generate.go && cat pool.go
