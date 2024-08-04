FROM golang:1.22-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src

RUN go build -o app ./src/app

FROM golang:1.22-bullseye

WORKDIR /root/

COPY --from=builder /app/app .

COPY src/config/config.yaml /root/config.yaml

EXPOSE 3000

CMD ["./app"]
