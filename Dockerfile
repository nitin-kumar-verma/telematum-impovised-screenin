FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


# Build the application
RUN go build -o main .

CMD [ "./main" ]

EXPOSE 8080