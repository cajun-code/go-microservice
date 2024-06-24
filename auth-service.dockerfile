FROM golang:1.22-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /app/dist/auth-service ./auth-service/cmd/api/

FROM alpine
RUN mkdir /app
COPY --from=builder /app/dist/auth-service /app
#COPY ./broker-service /app
CMD [ "/app/auth-service" ]