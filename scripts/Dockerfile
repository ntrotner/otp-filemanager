FROM golang:1.20.3

WORKDIR /usr/src/app

ENV HTTPPORT=8080

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN rm .env
RUN go build -v -o "otp-filemanager"
RUN mv otp-filemanager /usr/local/bin/otp-filemanager


EXPOSE 8080

CMD ["otp-filemanager"]