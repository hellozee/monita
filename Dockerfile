FROM golang:latest

RUN mkdir -p /go/src/monita
ENV GO111MODULE=on
WORKDIR /go/src/monita
COPY . /go/src/monita
RUN go install github.com/hellozee/monita
CMD /go/bin/monita

EXPOSE 8080