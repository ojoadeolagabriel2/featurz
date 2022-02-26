FROM golang:1.17.7-alpine3.15

WORKDIR $GOPATH/src/github.com/codefresh-contrib/go-sample-app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN rm -rf /go/pkg /go/src

EXPOSE 12345
CMD ["featurz"]