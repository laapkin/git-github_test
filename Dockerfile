FROM golang:1.25.3-alpine3.22

WORKDIR /opt/app

COPY ./main.go ./

CMD ["go", "run", "main.go"]
