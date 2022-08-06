FROM golang:1.18

WORKDIR /app

COPY . .

RUN go mod tidy && \
    go mod download && \ 
    go mod verify

RUN go build -v -o .

CMD ["go", "run", "main.go"]
