FROM golang

RUN apt-get update && apt-get install -y telnet

COPY ./ /app/

WORKDIR /app/

RUN go build .

EXPOSE 8080

CMD ["./front_api"]
