FROM alpine:3.14
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN echo "http://dl-4.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
RUN apk --update add --no-cache
RUN apk add --no-cache tcpreplay
RUN mkdir -p /pcap/
ADD ./two_account.pcap /pcap/
ADD ./two_account.cach /pcap/
