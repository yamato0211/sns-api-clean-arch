FROM golang:latest

COPY . /app/src
WORKDIR /app/src

RUN go install 

RUN go build -o /go/bin/air github.com/cosmtrek/air

CMD [ "go", "run", "main.go" ]