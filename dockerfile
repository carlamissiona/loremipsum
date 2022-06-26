FROM golang:1.18-bullseye

RUN go install github.com/carlamissiona/loremipsum

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/web
RUN mkdir -p "$APP_HOME"
RUN go build -o main .
WORKDIR "$APP_HOME"
EXPOSE 8010
CMD ["./main"]