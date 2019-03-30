FROM golang:1.10.4

RUN mkdir -p /go/src/GoItemAPI
WORKDIR /go/src/GoItemAPI

COPY *.go /go/src/GoItemAPI/
COPY *.json /go/src/GoItemAPI/

RUN go get
RUN go build

EXPOSE 8080

CMD ["./GoItemAPI"]