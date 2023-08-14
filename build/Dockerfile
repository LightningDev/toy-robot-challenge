FROM golang:1.20

WORKDIR /app

COPY .. .
RUN go mod download && go mod verify
RUN go build -o toy-robot .

ENTRYPOINT ["/app/toy-robot"]