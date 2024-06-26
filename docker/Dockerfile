# Start with Golang image to build src/
FROM golang:1.22.2-alpine AS builder

COPY src/ /src/
WORKDIR /src

RUN go mod tidy
RUN go build .

FROM alpine:latest
COPY --from=builder /src/echo-app /bin/echo-app
EXPOSE 8889
ENTRYPOINT ["/bin/echo-app"]