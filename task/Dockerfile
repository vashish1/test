FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go build -o main .
EXPOSE 3000
CMD ["/app/main"]