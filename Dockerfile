FROM golang:1.21.1-bullseye

RUN go version
ENV GOPATH=/

COPY ./ ./

#wait for postgres
RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

#run main
RUN go mod download
RUN go build -o marketplace ./cmd/main.go

CMD ["./marketplace"]