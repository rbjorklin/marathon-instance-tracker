FROM gliderlabs/alpine:latest
#FROM debian:latest
MAINTAINER robin.bjorklin@gmail.com
EXPOSE 9090

ADD instance-tracker-alpine /bin/instance-tracker

ENTRYPOINT ["/bin/instance-tracker"]
