# 从基础镜像安装依赖
FROM python:3.8-alpine AS build
# RUN mkdir /install
# WORKDIR /install
COPY requirements.txt /requirements.txt
COPY /etc/supervisord.conf /etc/supervisord.conf
COPY /tmp /tmp
COPY /shell ./shell
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.cloud.tencent.com/g' /etc/apk/repositories
RUN apk add --no-cache gcc make libc-dev libffi-dev
RUN pip install --no-cache-dir -i http://mirrors.cloud.tencent.com/pypi/simple/ --trusted-host mirrors.cloud.tencent.com -r /requirements.txt
RUN mkdir /ytoolsbox-req-custom
COPY ./ /ytoolsbox-req-custom
WORKDIR /ytoolsbox-req-custom
# ENTRYPOINT ["supervisord", "-c", "/etc/supervisord.conf"]