FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev libc-dev
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o server ./cmd/server


FROM alpine:3.22
RUN apk --no-cache add ca-certificates

WORKDIR /app


COPY --from=builder /app/server .
RUN mkdir -p /app/data
EXPOSE 8080
CMD ["./server"]
