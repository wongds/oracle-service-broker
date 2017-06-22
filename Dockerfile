FROM golang:1.7.5

MAINTAINER zhaozy@neunn.com

RUN mkdir -p /go/src/oracle-service-broker/

COPY . /go/src/oracle-service-broker/
WORKDIR /go/src/oracle-service-broker/

RUN cp -r /go/src/oracle-service-broker/vendor/* /go/src/ && \
    chmod 755 /go/src/oracle-service-broker/docker/run.sh && \
    go install -a -v github.com/beego/bee

EXPOSE 8000

ENTRYPOINT ["/go/src/oracle-service-broker/docker/run.sh"]