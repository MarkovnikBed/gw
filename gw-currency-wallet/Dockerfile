FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./my_app ./cmd/
 
FROM scratch

COPY --from=builder /app/my_app /my_app
COPY --from=builder /app/config.env /config.env
ENTRYPOINT [ "/my_app" ]
