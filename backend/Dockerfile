FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY ./ ./

RUN pwd
RUN ls
RUN go mod tidy
RUN go build -o ./bin/apptiket ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/.env ./
COPY --from=builder /app/bin/apptiket ./

EXPOSE 7010

RUN ls
CMD ./apptiket