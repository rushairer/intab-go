FROM golang:latest

WORKDIR /go/src/intab
COPY . .

RUN go-wrapper download
RUN go-wrapper install

CMD [ "go-wrapper", "run" ]
