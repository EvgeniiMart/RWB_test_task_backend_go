FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o liveTop ./cmd/liveTop



FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update && apt-get install -y curl

COPY --from=builder /app/liveTop .
COPY --from=builder /app/data ./data

RUN chmod +x ./liveTop
    
CMD ["./liveTop"]
