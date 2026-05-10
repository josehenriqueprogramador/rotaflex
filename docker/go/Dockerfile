FROM golang:1.25-alpine

WORKDIR /app

RUN apk add --no-cache git

COPY . .

RUN go mod tidy
RUN go mod download

RUN go build -o app .

EXPOSE 8000

CMD ["./app"]

