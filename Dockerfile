FROM golang:latest

LABEL maintaner="rafaeldoering <rafael.doering.soares@gmail.com>"

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

ENV PORT=:8000

RUN go build

CMD [ "./gopher-bank" ]
