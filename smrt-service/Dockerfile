
FROM golang:latest as builder

WORKDIR /go/src/smrt-service

COPY ./smrt-service .

RUN go get github.com/go-kit/kit/endpoint
RUN go get github.com/go-kit/kit/log
RUN go get github.com/go-kit/kit/transport/http

RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest


RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/smrt-service .

EXPOSE 8080
CMD ["./smrt-service"]