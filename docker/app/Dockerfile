FROM golang:1.21.1-alpine

WORKDIR /app

COPY ./codes/go.mod ./codes/go.sum ./
RUN go mod download && go mod verify

COPY ./codes/ ./

RUN go build -o /golang-binary

EXPOSE 8080

CMD [ "/golang-binary" ]
