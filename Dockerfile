FROM golang:latest

RUN mkdir /build

WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/Studio56School/university/cmd
RUN cd /build && git clone https://github.com/Studio56School/university.git

RUN cd /build/university/cmd && go build -o main

EXPOSE 8080

ENTRYPOINT ["/build/university/cmd/main"]