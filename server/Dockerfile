FROM golang:1.14

COPY . /home/GitHelper/server/

WORKDIR /home/GitHelper/server/

RUN go mod download -x

RUN go build -o server server.go

CMD ["./server"]
