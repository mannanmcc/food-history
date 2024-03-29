#/food-history/Dockerfile
FROM "golang:alpine"
# Here the FROM clause states which base image we are intending to work with. If the image does not exist, locally, Docker automatically fetches it from Dockerhub. If you supplied a URI for the image, Docker will download it from there too. Here we begin with the golang:alpine image

MAINTAINER "Abdul Mannan <mannanmcc@gmail.com>"
# Shows who created/maintains the file

WORKDIR /go/src
# Tells Docker to create a working directory that the container will by default use for your project. When you docker -ti <image> it will check into this folder first
COPY . /go/src
# This command tells Docker to copy files from our local machine, into the container that is being built. In some cases we can choose to download our code from Github or any other source, for our case it is simple enough to just COPY the food-history to Docker. and place it in the WORKDIR mentioned earlier.

RUN cd /go/src && go build -o main
# Check into our working directory, build our main.go

EXPOSE 8080
# This tells Docker to expose a certain port that can be listened to. This is important since our application is exposing on Port 8080, we need Docker to also expose on this port so that external sources can interact with our app.

ENTRYPOINT "./main"
# This is the first command to run once the container starts, turning it into an automatically running welcome-app server.