# syntax=docker/dockerfile:1
FROM golang:1.20

LABEL authors="Sayed Hesham, Abdulla Al Juffairi, Ahmed Haider"
LABEL version="1.0"
LABEL description="This project covers the basic of Dockers."

WORKDIR /app1

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /bin/main.go ./main.go

EXPOSE 8080:8080

CMD [ "/bin/main.go" ]
