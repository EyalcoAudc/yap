FROM golang:1.12-buster

# Install ca-certificates
RUN apt update && apt install -y bzip2 ca-certificates

# Update CA certificates
COPY audiocdes2016.crt /usr/local/share/ca-certificates/
RUN update-ca-certificates

RUN mkdir -p /yap/src
COPY . /yap/src/yap

ENV GOPATH=/yap
WORKDIR /yap/src/yap

RUN bunzip2 data/*.bz2

RUN go mod init
RUN go get .
RUN  go mod vendor
RUN go build .

EXPOSE 8000

ENTRYPOINT ["/yap/src/yap/yap", "api"]