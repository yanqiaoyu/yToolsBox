# yToolsBox

[README](README_en.md) | [中文文档](README.md)

yToolsBox 是一个 All-In-One 的 工具收纳与调度平台。目前支持收纳调度脚本(*.py *.sh)与 docker 容器。



* [yToolsBox](#ytoolsbox)
   * [为什么开发 yToolsBox ？](#为什么开发-ytoolsbox-)
   * [关键价值](#关键价值)
   * [部署](#部署)
      * [手动部署](#手动部署)
         * [1. 安装](#1-安装)
         * [2. 验证安装结果](#2-验证安装结果)
      * [用 docker-compose 部署](#用-docker-compose-部署)
         * [1.安装](#1安装)
         * [2.验证安装结果](#2验证安装结果)
   * [使用教程](#使用教程)
      * [1. 执行脚本类工具](#1-执行脚本类工具)
         * [1.1 准备好一个脚本文件](#11-准备好一个脚本文件)
         * [1.2 新增一个脚本工具](#12-新增一个脚本工具)
         * [1.3 填写工具基础信息](#13-填写工具基础信息)
         * [1.4 填写工具配置信息](#14-填写工具配置信息)
         * [1.5 确认添加](#15-确认添加)
         * [1.6 发起执行任务](#16-发起执行任务)
         * [1.7 查看任务执行结果](#17-查看任务执行结果)
      * [2. 执行容器类工具](#2-执行容器类工具)
         * [2.1 新增一个容器工具](#21-新增一个容器工具)
         * [2.2 填写容器工具的信息](#22-填写容器工具的信息)
         * [2.3 填写工具配置信息](#23-填写工具配置信息)
         * [2.4 确认添加](#24-确认添加)
         * [2.5 发起执行任务](#25-发起执行任务)
         * [2.6 查看任务执行结果](#26-查看任务执行结果)
      * [3. 为已有的工具，添加新的配置](#3-为已有的工具添加新的配置)
         * [3.1 进入某个工具界面](#31-进入某个工具界面)
         * [3.2 新增配置](#32-新增配置)
         * [3.3 新增任务](#33-新增任务)
   * [开发进度](#开发进度)



## 为什么开发 yToolsBox ？

2021年中下旬,我参与到了(**)项目当中,承担相应的测试与管理的工作. 

(**)作为一个非常成熟的产品,时任的开发与测试留存有许多能效改进的工具. 

当时我遇到了一个需求: 我需要对比不同版本(**)之间,关键配置的差异

在我向时任测试人员求助是否有现成的工具后,他开始翻箱倒柜地寻找有关(**)的文件对比工具,找到之后,他对我说:"你可以自己看下这个工具怎么用, 我也记不太清了".

在阅读完脚本如何实现,并根据实际情况传参后,这个关键配置对比的需求完成了.

执行脚本并获取结果大约只会花费10秒不到的时间,但我却把大多数的时间花费在:

1. 找人,找到有对应工具的人
2. 沟通,寻求工具的使用方法
3. 看源码,如果工具经历了好几手,没有什么人知道怎么用,那么就要自己看源码寻求使用方法

事实上,不仅(**)会遇到这个问题,所有经历过多次交接的项目都会有类似场景

所以我开始想:
	
**如果大多数的工具都是 执行+传参 这种模式,那么我们能不能统一封装起来,优化一下 找人+沟通+看源码 的过程, 提升工具复用的能力和效率呢?**

基于如上的背景, 我开始着手开发一套可以完成不同类型工具的收纳与调度的PaaS化解决方案

## 关键价值

* 集中管理我的脚本工具以及容器化后的工具

工作中总会有一些写的辅助工具, 一段时间没有使用后, 要么忘了位置, 要么忘了怎么用了

有了这个工具盒, 我终于可以不用担心这些问题了~

* 更加方便快捷的调用工具

写出来的工具或者脚本, 有时候会需要交付别人使用

通常情况下, 你需要告诉别人需要传入什么参数, 或者给别人一些文档, 告诉他们如何使用

把工具集成到工具盒中后, 调用工具需要选择好, 你添加这个工具时的配置模板, 或者根据配置模板新建一个自己的配置, 就能直接开始使用了

交付效率一下就提升了~

* 支持定时调用工具

在做稳定性或者压力测试的时候, 可能会出现一些定时校验或者定时加条件的场景

所以这个工具盒除了普通形式的调用工具, 也支持添加定时任务, 让工具盒自己去调用工具

## 部署

### 手动部署

#### 1. 安装

依次执行如下指令

```shell
docker network create --driver bridge ytoolsbox_network

docker volume create ytoolsbox_db-data

docker run -itd -p 5432:5432 --name yToolsBox-db --network ytoolsbox_network -e POSTGRES_PASSWORD=`YOUR_PASSWORD` -v ytoolsbox_db-data:/var/lib/postgresql/data postgres

docker run -itd -p 8081:8081 --name yToolsBox-api --network ytoolsbox_network -e HOST_SCRIPT_PATH=/home/yToolsBox/api/Script -v /home/yToolsBox/api/Script:/root/Script yanqiaoyu/ytoolsbox-api:v0.2.3  supervisord -c /etc/supervisord.conf

docker run -itd -p 80:80 --network ytoolsbox_network --name yToolsBox-dashboard yanqiaoyu/ytoolsbox-dashboard:v0.2.3
```

#### 2. 验证安装结果

执行docker ps出现以下结果，则证明镜像拉起成功

![manu_deploy](/doc/pic/manu_deploy1.png)

访问 http://yourIP 验证是否安装成功

### 用 docker-compose 部署

#### 1. 安装

```shell
注：使用docker-compose的安装方式会现场编译一次镜像，若环境中缺少相关的依赖，不建议采用这种方式安装
```

```shell
git clone https://github.com/yanqiaoyu/yToolsBox.git
cd yToolsBox
docker-compose up -d
```

#### 2.验证安装结果

同上

## 使用教程

### 1. 执行脚本类工具

#### 1.1 准备好一个脚本文件

假设我们现在有一个脚本 pwd.sh，具体内容如下

![usage_script_1](/doc/pic/usage_script_1.png)

#### 1.2 新增一个脚本工具

进入新增工具盒的界面

![usage_script_2](/doc/pic/usage_script_2.png)

#### 1.3 填写工具基础信息

工具基础信息中，仅工具名称与作者名称为必填选项，其余内容选填（工具简介建议填写）

![usage_script_3](/doc/pic/usage_script_3.png)

#### 1.4 填写工具配置信息

工具配置信息中，我们先看上半部分

![usage_script_4](/doc/pic/usage_script_4.png)

* 工具类型

这里我们选择脚本工具

* 上传文件

这里我们上传我们自己写好的脚本文件

* 脚本名称

脚本名称会根据上传的文件自动识别，无需填写

* Shell版本

因为我们上传的是一个shell脚本，因此这里会需要选择shell解释器版本

考虑到一些旧的操作系统，或者经过裁剪的操作系统，可能不兼容bash，只兼容sh，所以设置了这样一个选项

```shell
注：如果上传的是Python文件，那么这里需要选择的就是Python版本
```

* 脚本绝对路径

因为我们在下面选择了远程执行，所以这里需要填写脚本存放在远程服务器的路径

```shell
注：对于需要远程执行的脚本工具，平台会先将脚本上传到我们所指定的脚本绝对路径下，然后再ssh进远程环境执行
```

* 运行参数

这里我们填写的是 /home , 这意味着接下来执行时，会去列出远程环境中/home路径下的所有文件

当然，不是所有的工具都需要参数，所以这个参数不是必填的

* 最终执行语句

最终执行语句也无需填写，直接根据我们上面所填写的所有信息自动生成

* 本地执行or远程执行

这里我们选择远程执行

```shell
注：
如果选择的是本地执行，那么脚本会在平台所安装的宿主机被执行
与远程执行类似，平台会ssh进宿主机，进入存放脚本的位置执行脚本
宿主机存放在宿主机的 /home/yToolsBox/api/Script/`工具名称`/ 位置下
```

再看到下半部分

![usage_script_5](/doc/pic/usage_script_5.png)

* SSH信息

按照实际情况，填写SSH的IP，端口，账号，密码即可

填写完成后，点击下一步即可

#### 1.5 确认添加

确认一切信息正常后，提交即可

![usage_script_6](/doc/pic/usage_script_6.png)

#### 1.6 发起执行任务

新增好工具之后，进入任务界面，新增任务

![usage_script_7](/doc/pic/usage_script_7.png)

![usage_script_8](/doc/pic/usage_script_8.png)

![usage_script_9](/doc/pic/usage_script_9.png)


#### 1.7 查看任务执行结果

等待任务执行完成后，我们可以获取任务的执行结果

![usage_script_10](/doc/pic/usage_script_10.png)

![usage_script_12](/doc/pic/usage_script_12.png)

### 2. 执行容器类工具

容器工具的调度执行更为简单，因为大部分的工作docker都帮我们做完了，下面以执行一次docker官方的hello-world容器为例，不多做言语赘述

#### 2.1 新增一个容器工具

同上

#### 2.2 填写容器工具的信息

![usage_script_21](/doc/pic/usage_script_21.png)

#### 2.3 填写工具配置信息

```
注：你需要保证执行docker工具的虚拟机docker相关的功能是正常的
```

![usage_script_22](/doc/pic/usage_script_22.png)

![usage_script_23](/doc/pic/usage_script_23.png)

#### 2.4 确认添加

同上

#### 2.5 发起执行任务

同上

#### 2.6 查看任务执行结果

![usage_script_24](/doc/pic/usage_script_24.png)

### 3. 为已有的工具，添加新的配置

我们希望针对不同的环境或者对象复用我们的工具，这种情况就可以为已有的工具添加一个新的配置，达到工具快速复用的目的

#### 3.1 进入某个工具界面

![usage_script_31](/doc/pic/usage_script_31.png)

#### 3.2 新增配置

![usage_script_32](/doc/pic/usage_script_32.png)

这里配置名称与配置简介是必填的

可以看到，运行参数填写的是/ ，这样我们利用用一个工具，去检查根目录下的文件

![usage_script_33](/doc/pic/usage_script_33.png)

填写好之后，确认新增

![usage_script_34](/doc/pic/usage_script_34.png)

#### 3.3 新增任务

添加好配置之后，我们进入新增任务界面，可以看到已经添加的配置

![usage_script_35](/doc/pic/usage_script_35.png)

选中这个配置，发起一次任务，观察执行结果

![usage_script_36](/doc/pic/usage_script_36.png)

## 开发进度

截止至 2021年10月25日23:30:37

[需求矩阵与开发进度](https://note.youdao.com/s/6pur0oT7)

![developProgress](/doc/pic/developProgress.png)
