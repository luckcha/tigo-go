# Step 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o tigo-go

# Step 2: Run
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/tigo-go .
EXPOSE 10000
CMD ["./tigo-go"]
