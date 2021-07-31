FROM golang:latest

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
COPY config.yaml ./
COPY config ./config
COPY db ./db
COPY issue ./issue
COPY push ./push

RUN go build -o /github-webhook-server

EXPOSE 8080

CMD [ "/github-webhook-server" ]
