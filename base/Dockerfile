FROM golang:1.7.5

MAINTAINER wang-x@neunn.com

COPY . /root

ENV PKG_CONFIG_PATH /root
RUN apt-get update && \
    apt-get install -y libaio-dev && \
    apt-get install -y unzip && \
    rm -rf instantclient-basic-linux.x64-11.2.0.3.0.zip && \
    rm -rf instantclient-sdk-linux.x64-11.2.0.3.0.zip
RUN ln -s /root/instantclient_11_2/libclntsh.so.11.1 /root/instantclient_11_2/libclntsh.so && \
    mv /root/ld.so.conf /etc/ld.so.conf && \
    ldconfig
