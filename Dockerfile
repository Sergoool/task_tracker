FROM golang:1.25-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o task_tracker ./cmd/app/main.gi

CMD [ "./task_tracker" ]


