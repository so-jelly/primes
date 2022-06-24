FROM golang:1.18 as builder
WORKDIR /build

# Use docker cache for deps
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

# our source
COPY main.go main.go
COPY pkg/ pkg/
RUN go vet -v ./...
RUN go test -v ./...

# build
RUN CGO_ENABLED=0 go build -o /go/bin/app

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static:nonroot
COPY --from=builder /go/bin/app /
# start primes when container starts
ENTRYPOINT [ "/app" ]
# use "nonroot" user
USER 65532:65532
# a docker volume for our input data
VOLUME [ "/data" ]
