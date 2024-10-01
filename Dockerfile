FROM golang:1.23


RUN mkdir /app

WORKDIR /app

COPY . .

EXPOSE 8080

CMD ["go", "run", "main.go"]
