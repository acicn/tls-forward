FROM golang:1.14 AS builder
ENV CGO_ENABLED 0
WORKDIR /go/src/app
ADD . .
RUN go build -o /tls-forward

FROM scratch
COPY --from=builder /tls-forward /tls-forward
CMD ["/tls-forward"]