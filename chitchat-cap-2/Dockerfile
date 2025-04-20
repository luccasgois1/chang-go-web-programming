FROM golang:1.24.2

COPY ./data /app/data
COPY ./public /app/public
COPY ./templates /app/templates
COPY ./*.go /app/
COPY ./*.mod /app/
COPY ./*.sum /app/

WORKDIR /app

RUN go build -o chitchat

CMD ["./chitchat"]
