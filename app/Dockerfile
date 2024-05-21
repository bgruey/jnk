FROM golang:1.22

WORKDIR /app

COPY ./ ./

RUN go mod init test-docker
RUN go mod tidy
RUN go build -o /msgg main.go

CMD ["/msgg"]