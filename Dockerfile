FROM golang
WORKDIR /go/src/app
RUN go get github.com/mesg-foundation/core/api/service && \
    go get github.com/bluele/slack
COPY . .
RUN go install -v ./...
RUN go build main.go
CMD [ "./main" ]