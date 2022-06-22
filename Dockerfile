FROM golang:1.18 as builder
WORKDIR /build
COPY . .
RUN go build -o /app/primes

FROM scratch
COPY --from=builder /app /app
ENTRYPOINT [ "/app/primes" ]
