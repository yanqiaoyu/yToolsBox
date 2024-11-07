FROM alpine:3.20.3
#RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
#RUN echo "http://dl-4.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
#RUN apk --update add --no-cache
# 使用 edge 的 testing 仓库来安装 tcpreplay
RUN apk add --no-cache tcpreplay --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing
RUN mkdir -p /pcap/
ADD ./InnerIPAll.pcap /pcap/
ADD ./ForiegnIPAll.pcap /pcap/
# ADD ./two_account.pcap /pcap/
# ADD ./two_account.cach /pcap/