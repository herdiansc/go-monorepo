#build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /auth-svc

#final stage
FROM alpine:3.21

LABEL Name="auth-svc" Version="1.0"

WORKDIR /root/

COPY --from=builder /auth-svc /root/
COPY --from=builder /app/.env.example /root/.env

EXPOSE 9001/tcp

CMD ["/root/auth-svc"]
