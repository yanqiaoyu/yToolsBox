# yToolsBox

[README](README_en.md) | [中文文档](README.md)

yToolsBox is an All-In-One platform which is for tool storage and scheduling。It currently supports storing and scheduling scripts (*.py *.sh) and containers.

## Why would I develop this yToolsBox ？

The emergence of shared bicycles has largely solved the "last mile" problem of people returning home from get off work.

We will encounter many similar scenarios in our daily work，for example：In order to improve daily work efficiency，we usually would produce some scripts or docker images to help us。But usually no one knows how to use each other’s tools；Or there are related documents for use, but few people are willing to read it。It usually ended hurriedly with "Could you please help me doing this？"

If we regard "solving the problems encountered in our respective work" as "going home", and taking "the various tools produced to solve these problems" as the "means to go home", then it is not difficult to find that this is actually another kind of "last mile" problem.

Therefore, the significance of yToolsBox is that it is a platform for storing and dispatching "transportation tools". Anyone can add the "transportation tools" produced by themselves to this platform in a specified way, and let others follow if necessary. Reuse, then go directly to the tool box to select the configuration and execute it, saving a lot of communication costs

## Deploy

### Manually deploy

#### 1. Install

Excute the following cmds

```shell
docker network create --driver bridge ytoolsbox_network

docker volume create ytoolsbox_db-data

docker run -itd -p 5432:5432 --name yToolsBox-db --network ytoolsbox_network -e POSTGRES_PASSWORD=test123456 -v ytoolsbox_db-data:/var/lib/postgresql/data postgres

docker run -itd -p 8081:8081 --name yToolsBox-api --network ytoolsbox_network -e HOST_SCRIPT_PATH=/home/yToolsBox/api/Script -v /home/yToolsBox/api/Script:/root/Script yanqiaoyu/ytoolsbox-api:v0.1.1  ./main -m production

docker run -itd -p 80:80 --network ytoolsbox_network --name yToolsBox-dashboard yanqiaoyu/ytoolsbox-dashboard:v0.1.1
```

#### 2. Verify the result

Execute "docker ps" to verify the status of the containers

![manu_deploy](/doc/pic/manu_deploy1.png)

Visit http://yourIP to verify whether the frontend has been deployed

### Use docker-compose

#### 1. Install

```shell
Notice：It's not recommended to use docker-compose if your envirement lack of dependencies which are used for compiling docker images 
```

```shell
git clone https://github.com/yanqiaoyu/yToolsBox.git
cd yToolsBox
docker-compose up -d
```

#### 2. Verify the result

Same  as above

## Tutorial

(TBD)

## Develop Progress

As of 23:30:37, October 25, 2021
![developProgress](/doc/pic/developProgress.png)
