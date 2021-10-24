FROM golang:1.16-alpine as builder
WORKDIR /pesakit
COPY pesakit .
RUN go clean -modcache
RUN go mod tidy
RUN go mod download
RUN cd cmd && go build -o pesakit

FROM alpine:latest
LABEL name="pesakit - commandline tool for mobile payment api"
RUN apk --no-cache add ca-certificates
WORKDIR /pesakit/
COPY --from=builder /pesakit ./
ENTRYPOINT ["cmd/pesakit"]
CMD ["--help"]