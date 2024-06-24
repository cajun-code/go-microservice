#FROM golang:1.22-alpine AS builder
# RUN mkdir /app
# WORKDIR /app
# COPY . .
# RUN go mod download

# RUN go build -o /app/broker-service ./broker-service/cmd/api/

FROM alpine
RUN mkdir /app
#COPY --from=builder /app/broker-service /app
COPY ./broker-service /app
CMD [ "/app/broker-service" ]