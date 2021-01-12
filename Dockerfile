FROM golang:1.15.6

RUN apt update && \
    apt-get -y install vim

