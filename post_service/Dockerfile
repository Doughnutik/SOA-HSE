FROM golang:1.23.7
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /post_service
EXPOSE 8082
ENTRYPOINT ["/post_service"]