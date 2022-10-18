FROM golang

RUN apt-get update && apt-get install -y telnet

COPY ./ /app/

WORKDIR /app/

RUN go build .

EXPOSE 8081

CMD ["./server"]
