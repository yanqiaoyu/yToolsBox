# yToolsBox

[README](README_en.md) | [中文文档](README.md)

yToolsBox is an All-In-One platform which is for tool storage and scheduling。It currently supports storing and scheduling scripts (_.py _.sh) and containers.

## Why would I develop this yToolsBox ？

The emergence of shared bicycles has largely solved the "last mile" problem of people returning home from get off work.

We will encounter many similar scenarios in our daily work，for example：In order to improve daily work efficiency，we usually would produce some scripts or docker images to help us。But usually no one knows how to use each other’s tools；Or there are related documents for use, but few people are willing to read it。It usually ended hurriedly with "Could you please help me doing this？"

If we regard "solving the problems encountered in our respective work" as "going home", and taking "the various tools produced to solve these problems" as the "means to go home", then it is not difficult to find that this is actually another kind of "last mile" problem.

Therefore, the significance of yToolsBox is that it is a platform for storing and dispatching "transportation tools". Anyone can add the "transportation tools" produced by themselves to this platform in a specified way, and let others follow if necessary. Reuse, then go directly to the tool box to select the configuration and execute it, saving a lot of communication costs

## Deploy

### Manually deploy

#### 1. Excute the following cmds

```shell
docker network create --driver bridge ytoolsbox_network

docker volume create ytoolsbox_db-data

docker run -itd -p 5432:5432 --name yToolsBox-db --network ytoolsbox_network -e POSTGRES_PASSWORD=test123456 -v ytoolsbox_db-data:/var/lib/postgresql/data postgres

docker run -itd -p 8081:8081 --name yToolsBox-api --network ytoolsbox_network -e HOST_SCRIPT_PATH=/home/yToolsBox/api/Script -v /home/yToolsBox/api/Script:/root/Script yanqiaoyu/ytoolsbox-api:v0.1.1  ./main -m production

docker run -itd -p 80:80 --network ytoolsbox_network --name yToolsBox-dashboard yanqiaoyu/ytoolsbox-dashboard:v0.1.1
```

#### 2. The cmds result

![manu_deploy](/doc/pic/manu_deploy1.png)

#### 3. Testing

Visit http://yourIP to verify whether the platform has been deployed

### Use docker-compose

#### 1. docker-compose.yml

```yaml
  version: '3'
  services:
    yToolsBox-db:
      container_name: 'yToolsBox-db'
      image: postgres
      restart: always
      ports:
        - 5432:5432
      volumes:
        - db-data:/var/lib/postgresql/data
      networks:
        - network
      environment:
        POSTGRES_PASSWORD: test123456

    yToolsBox-api:
      container_name: 'yToolsBox-api'
      build:
        context: ./ytoolsbox-gin
        dockerfile: Dockerfile
      image: yanqiaoyu/ytoolsbox-api:v0.1.1
      depends_on:
        - yToolsBox-db
      networks:
        - network
      volumes:
        - /home/yToolsBox/api/Script:/root/Script
      # should be same as above path
      environment:
        HOST_SCRIPT_PATH: /home/yToolsBox/api/Script
      ports:
        - 8081:8081
      command: ["sh", "wait-for", "yToolsBox-db:5432", "--", "./main", "-m", "production"]

    yToolsBox-dashboard:
      container_name: 'yToolsBox-dashboard'
      build:
        context: ./ytoolsbox-vue
        dockerfile: Dockerfile
      image: yanqiaoyu/ytoolsbox-dashboard:v0.1.1
      networks:
        - network
      ports:
        - 80:80
      depends_on:
        - yToolsBox-db
        - yToolsBox-api

  volumes:
    db-data:
  networks:
    network:
      driver: bridge
```

#### 2.Deploy

Excute the following cmds

```shell
docker-compose up -d
```

#### 3.Testing

Same things as above

## Tutorial

(TBD)

## Develop Progress

As of 15:06:52, October 8, 2021
![developProgress](/doc/pic/developProgress.png)
