FROM golang:1.22-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /app/dist/logger-service ./logger-service/cmd/api/

FROM alpine
RUN mkdir /app
COPY --from=builder /app/dist/logger-service /app
#COPY ./logger-service /app
CMD [ "/app/logger-service" ]