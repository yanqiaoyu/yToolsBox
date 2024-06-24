FROM nginx:alpine
COPY dist/ /usr/share/nginx/html/
COPY nginx/ /etc/nginx/
# # 安装依赖
# RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/proxy.mirrors.sangfor.org/g' /etc/apk/repositories \
#     && apk add --no-cache git \
#     && apk add --no-cache --virtual .gyp python make g++ \
#     && npm install \
#     && apk del .gyp

# # 构建 Vue 项目
# RUN npm run build
