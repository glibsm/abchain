FROM docker.io/library/golang:1.20.5-alpine3.18 as builder
WORKDIR /go/src/github.com/glibsm/abchain
COPY . .
RUN CGO_ENABLED=0 go build .

FROM docker.io/library/alpine:latest
COPY --from=builder /go/src/github.com/glibsm/abchain/abchain /usr/bin
CMD ["/usr/bin/abchain"]
