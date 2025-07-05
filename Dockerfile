FROM golang:1.24-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/air-verse/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/a-h/templ/cmd/templ@latest

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]