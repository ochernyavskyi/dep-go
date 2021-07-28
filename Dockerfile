FROM golang

WORKDIR /go/src/app
COPY my-app.go .
RUN go build my-app.go
EXPOSE 80
CMD ["./my-app"]