FROM golang:1.22

WORKDIR /usr/src/todo-app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download 

COPY . .
RUN go build -o /usr/local/bin/todo-app ./cmd/main.go

CMD ["todo-app"]