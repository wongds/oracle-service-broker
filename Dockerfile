FROM neunnsy/oracle-service-broker:11.2

MAINTAINER zhaozy@neunn.com

RUN mkdir -p /go/src/github.com/compassorg/oracle-service-broker/

COPY . /go/src/github.com/compassorg/oracle-service-broker/
WORKDIR /go/src/github.com/compassorg/oracle-service-broker/

RUN cp -r /go/src/github.com/compassorg/oracle-service-broker/vendor/* /go/src/ && \
    chmod 755 /go/src/github.com/compassorg/oracle-service-broker/docker/run.sh && \
    go install -a -v github.com/beego/bee

ENV PKG_CONFIG_PATH /root

EXPOSE 8000

ENTRYPOINT ["/go/src/github.com/compassorg/oracle-service-broker/docker/run.sh"]