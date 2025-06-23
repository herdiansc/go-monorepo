#build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /logistic-svc

#final stage
FROM alpine:3.21

LABEL Name="logistic-svc" Version="1.0"

WORKDIR /root/

COPY --from=builder /logistic-svc /root/
COPY --from=builder /app/.env.example /root/.env

EXPOSE 9001/tcp

CMD ["/root/logistic-svc"]
