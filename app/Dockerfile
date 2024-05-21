FROM golang:1.22

WORKDIR /app

COPY ./ ./
ARG GONOPROXY=github.com

RUN rm go.mod 
RUN go mod init test-docker
RUN go mod tidy
RUN go build -o /msgg main.go

CMD ["/msgg"]