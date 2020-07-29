FROM golang:1.14

WORKDIR /go/src/app

ENV DB_USER="your-db-user" DB_PASSWORD="your-db-password" DB_PORT=5432 \
 DB_NAME="your-db-name" DB_HOST="localhost"

COPY . .

RUN go mod download && go install ./... && go build -o main .

EXPOSE 8081

CMD ["./main"]
