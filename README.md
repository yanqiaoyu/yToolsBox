# yToolsBox
[README](README_en.md) | [中文文档](README.md)

yToolsBox是一个小型的工具收纳与调度平台。目前支持收纳调度脚本(*.py *.sh)与容器。

## 为什么开发 yToolsBox ？

共享单车的出现，很大程度上解决了人们下班回家的“最后一公里”问题。

同样的，日常的工作中我们也会遇到很多与之类似的场景，比方说：我们会为了提升日常的工作效率，产出各种辅助工作的脚本或者docker镜像。但往往每个人产出的辅助工具都处于一种“各自为战”的割裂状态，大家互相都不了解对方的工具如何使用；亦或者有相关的使用说明文档，却鲜有人愿意阅读，最终还是以“还是你帮我弄一下吧”草草收场

如果我们把“解决各自工作中遇到的问题”当做“回家”，把为了解决这些问题而“产出的形形色色的工具”当做“回家的手段”，那么不难发现，这其实也是另一种形式的“最后一公里”问题。

因此，yToolsBox的意义就在于，它是一个“交通工具”的收纳与调度平台，任何人都可以把自己产出的“交通工具”，按照指定的方式添加进这个平台，后续如果需要让其他人复用，那么直接进入工具盒选择配置并执行即可，节省了大量的沟通成本

## 部署

### 手动部署

#### 1. 依次执行如下指令

```shell
docker network create --driver bridge yToolsBox-network

docker volume create yToolsBox-db-data

docker run -itd --name yToolsBox-db --network yToolsBox-network -e POSTGRES_PASSWORD=test123456 -v yToolsBox-db-data:/var/lib/postgresql/data postgres

docker run -itd --name yToolsBox-api --network yToolsBox-network -v /home/yToolsBox/api/Script:/root/Script yanqiaoyu/ytoolsbox-api:v0.1

docker run -itd -p 80:80 --network yToolsBox-network --name yToolsBox-dashboard yanqiaoyu/ytoolsbox-dashboard:v0.1
```

#### 2. 部署结果

![manu_deploy](/doc/pic/manu_deploy1.png)

#### 3. 验证结果

访问 http://yourIP 验证是否安装成功

### 用docker compose部署

#### 1. docker-compose

```yaml
  version: '3'
  services:
    yToolsBox-db:
      #容器名称
      container_name: 'yToolsBox-db'
      # 镜像名称
      image: postgres
      # docker重启的时候会自动重启容器
      restart: always
      # 内部端口映射到宿主机
      # ports:
      #   - 5432:5432
      expose:
        - "5432"
      # 挂载
      volumes:
        - db-data:/var/lib/postgresql/data
      # 加入网络
      networks:
        - network
      # 设置环境变量
      environment:
        POSTGRES_PASSWORD: test123456

    yToolsBox-api:
      #容器名称
      container_name: 'yToolsBox-api'
      # 镜像名称
      image: yanqiaoyu/ytoolsbox-api:v0.1
      # 控制启动顺序
      depends_on:
        - yToolsBox-db
      # 加入网络
      networks:
        - network
      # 挂载
      volumes:
        - /home/yToolsBox/api/Script:/root/Script
      # 允许docker网络中其他容器访问，但是不允许外部访问
      # ports:
      #   - 8081:8081
      expose:
        - "8081"


    yToolsBox-dashboard:
      #容器名称
      container_name: 'yToolsBox-dashboard'
      # 镜像名称
      image: yanqiaoyu/ytoolsbox-dashboard:v0.1
      # 加入网络
      networks:
        - network
      ports:
        - 80:80
      # 控制启动顺序
      depends_on:
        - yToolsBox-db
        - yToolsBox-api

  # 创建卷
  volumes:
    db-data:
  # 创建网络
  networks:
    network:
      driver: bridge

```

#### 2.安装

执行
```shell
docker-compose up -d
```



## 使用教程

(针对脚本与docker镜像2种类型都进行演示，TBD)

## 开发进度

截止至 2021年10月8日15:06:52
![developProgress](/doc/pic/developProgress.png)