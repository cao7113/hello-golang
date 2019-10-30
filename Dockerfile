FROM golang:1.13-alpine3.10 AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
# Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    main.go

FROM alpine:3.10 AS final
COPY --from=builder /app/main /app/
WORKDIR /app
CMD ["./main"]
# EXPOSE 50051