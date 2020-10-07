FROM docker.io/library/golang:1.14.6-alpine3.12 as builder
WORKDIR /go/src/github.com/glibsm/abchain
COPY . .
RUN CGO_ENABLED=0 go build .

FROM docker.io/library/alpine:3.12
COPY --from=builder /go/src/github.com/glibsm/abchain/abchain /usr/bin
CMD ["/usr/bin/abchain"]
