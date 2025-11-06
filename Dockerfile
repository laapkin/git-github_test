    FROM golang:1.25.3-alpine3.22
    
    WORKDIR /app

    COPY ./main.go /app/

    CMD ["go", "run", "main.go"]
 


        

