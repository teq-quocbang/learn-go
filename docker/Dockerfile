FROM golang:1.22.6-alpine3.19

WORKDIR /

COPY . .

RUN go mod download

RUN go build -o ./learn main.go

CMD [ "/learn" ]
