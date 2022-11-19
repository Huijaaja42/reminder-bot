FROM golang:1.19

WORKDIR /usr/src/reminder-bot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ADD https://raw.githubusercontent.com/objectbox/objectbox-go/main/install.sh .
RUN chmod +x ./install.sh
RUN ./install.sh

RUN go build -v -o /usr/local/bin/reminder-bot .

CMD ["reminder-bot"]