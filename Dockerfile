FROM golang:1.22.6-alpine3.19 as builder

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o ./learn main.go

CMD [ "/learn" ]

FROM alpine:3.20.2

WORKDIR /server

COPY --from=builder /learn /server/learn

EXPOSE 8080

CMD [ "/server/learn" ]