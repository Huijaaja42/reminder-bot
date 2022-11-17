FROM golang:1.19

WORKDIR /usr/src/reminder-bot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/reminder-bot .

CMD ["reminder-bot"]