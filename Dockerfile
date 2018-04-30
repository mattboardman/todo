FROM golang:1.10.1 AS build-env
RUN go get -d -v github.com/google/uuid github.com/gorilla/handlers github.com/gorilla/mux
ADD . /go/src/ToDo
RUN cd /go/src/ToDo && CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build-env /go/src/ToDo/ .
ENTRYPOINT ./app
EXPOSE 1080 8080