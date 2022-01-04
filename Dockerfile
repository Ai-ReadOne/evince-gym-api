FROM golang:alpine

ENV PORT=8000

WORKDIR /evince-gym-api

COPY ./go.mod .
COPY ./go.sum .

RUN apk update && apk add --no-cache git
RUN go mod download

COPY . .

RUN go build

CMD ["./evince-gym-api"]

EXPOSE $PORT