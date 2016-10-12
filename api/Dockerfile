FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/KanybekMomukeyev/dockerCompose/api

# Install our dependencies
RUN go get github.com/go-martini/martini
RUN go get github.com/martini-contrib/binding
RUN go get github.com/martini-contrib/render

RUN go get gopkg.in/mgo.v2

# Install api binary globally within container
RUN go install github.com/KanybekMomukeyev/dockerCompose/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 3000