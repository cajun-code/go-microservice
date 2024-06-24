FROM golang:1.22-alpine AS builder
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o /app/dist/front-end ./front-end/cmd/web/

FROM alpine
RUN mkdir /app
COPY --from=builder /app/dist/front-end /app
COPY --from=builder /app/front-end/cmd/web/templates /app/cmd/web/templates 
#COPY ./broker-service /app
WORKDIR /app
CMD [ "./front-end" ]