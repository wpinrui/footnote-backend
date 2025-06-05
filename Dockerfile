FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping cmd/main.go

EXPOSE 8080

CMD ["/docker-gs-ping"]
