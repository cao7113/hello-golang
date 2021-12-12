FROM golang:1.15-alpine3.12 AS builder
RUN apk update && apk add --no-cache git

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
# Mark the build as statically linked.
RUN CGO_ENABLED=0 go build -v \
    -installsuffix 'static' \
    -o /app/bin/rpc \
    rpc/main.go

FROM alpine:3.12 AS final
WORKDIR /app
COPY --from=builder /app/bin /app/bin
ENV APP_ENV=production
CMD ["bin/rpc"]
EXPOSE 50051