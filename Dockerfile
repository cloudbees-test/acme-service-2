
# builder container
FROM golang:alpine as builder

WORKDIR /app 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/api-server .

# runtime container
FROM alpine:3.18.2

WORKDIR /app
COPY --from=builder /app/build/api-server /app/

EXPOSE 8080
ENTRYPOINT ["/app/api-server"]