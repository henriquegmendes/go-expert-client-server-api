FROM golang:1.19.4-buster

RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y sqlite3 libsqlite3-dev

WORKDIR /app

COPY . .

RUN ["go", "build", "-o", "main"]

EXPOSE 8000

CMD ["./main", "server"]