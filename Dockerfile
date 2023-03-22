FROM golang:1.19.0

WORKDIR /github.com/haritechversant/go-sample-docker

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy

CMD ["air", "./cmd/main.go", "-b", "0.0.0.0"]