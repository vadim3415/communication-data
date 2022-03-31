# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o cmd ./cmd/main.go && go build -o data_simulator ./data_simulator/main.go

COPY start.sh start.sh
CMD ["sh", "./start.sh"]

######## docker build --tag diplom .
######## docker run --rm -it diplom
######## docker run --rm -p 8006:8006 -it diplom
######## docker run --rm -p 8006:8006 -p 9002:9002 -it diplom
