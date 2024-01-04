FROM golang:latest

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY ./backend .

RUN go mod download

CMD ["go", "run", "main.go"]
