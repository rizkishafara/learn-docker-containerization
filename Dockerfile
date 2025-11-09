FROM golang:1.21-alpine
WORKDIR /app
COPY src ./src
WORKDIR /app/src
RUN go mod download
RUN go build -o ../main .
WORKDIR /app
EXPOSE 3000
CMD ["./main"]
