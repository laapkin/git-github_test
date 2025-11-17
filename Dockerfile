FROM golang:1.25.3-alpine3.22

WORKDIR /opt/app

COPY ./go.mod ./

COPY ./ ./

CMD ["go", "run", "./cmd/main.go"]
