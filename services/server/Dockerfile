FROM golang

COPY ./ /app/

WORKDIR /app/

RUN ["go", "build", "."]

CMD ["./server"]

EXPOSE 8081
