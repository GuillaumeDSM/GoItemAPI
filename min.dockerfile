FROM scratch

WORKDIR /go/bin

COPY GoItemAPI /go/bin/GoItemAPI

EXPOSE 8080

CMD ["./GoItemAPI"]