FROM neunnsy/oracle-service-broker:11.2

MAINTAINER zhaozy@neunn.com

RUN mkdir -p /go/src/oracle-service-broker/

COPY . /go/src/oracle-service-broker/
WORKDIR /go/src/oracle-service-broker/

RUN cp -r /go/src/oracle-service-broker/vendor/* /go/src/ && \
    chmod 755 /go/src/oracle-service-broker/docker/run.sh && \
    go install -a -v github.com/beego/bee

ENV PKG_CONFIG_PATH /root/instantclient_11_2

EXPOSE 8000

ENTRYPOINT ["/go/src/oracle-service-broker/docker/run.sh"]