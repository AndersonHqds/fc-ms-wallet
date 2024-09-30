FROM golang:1.23

WORKDIR /app

RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN cd cmd/walletcore && go build

CMD ["./cmd/walletcore/walletcore"]