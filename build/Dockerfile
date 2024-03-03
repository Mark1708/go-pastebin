FROM golang:alpine

WORKDIR /pastebin
COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/pastebin ./cmd/pastebin

CMD ["/pastebin/bin/pastebin"]
EXPOSE 8080