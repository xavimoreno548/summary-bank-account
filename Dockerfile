FROM --platform=linux/amd64 golang:1.19

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main ./cmd/api/main.go

EXPOSE 8080
CMD ["./main"]
