FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/mannanmcc/food-history/api

# Install our dependencies
# RUN go get github.com/go-martini/martini
RUN go get github.com/jinzhu/gorm
RUN go get github.com/gorilla/mux
RUN go get github.com/jinzhu/gorm/dialects/mysql

# Install api binary globally within container 
RUN go install github.com/mannanmcc/food-history/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 3000
