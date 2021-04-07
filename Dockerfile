FROM golang:1.14

WORKDIR /app
ADD ./ /app
RUN go install
RUN go build -o index

ENTRYPOINT ["/app/index"]

