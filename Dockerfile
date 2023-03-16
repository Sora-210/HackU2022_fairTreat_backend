FROM golang:1.20.1-buster

RUN apt-get update -y &&\
    apt-get upgrade -y &&\
    apt-get install zip -y &&\
    PB_REL="https://github.com/protocolbuffers/protobuf/releases" &&\
    curl -LO $PB_REL/download/v22.0/protoc-22.0-linux-x86_64.zip &&\
    unzip protoc-22.0-linux-x86_64.zip -d $HOME/.local &&\
    export PATH="$PATH:$HOME/.local/bin"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 &&\
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 &&\
    export PATH="$PATH:$(go env GOPATH)/bin"

RUN mkdir /app

COPY . /app
WORKDIR /app/fairtreat/
RUN go install
CMD go run main.go ./grpc
EXPOSE 50000
