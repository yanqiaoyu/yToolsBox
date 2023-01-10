# 数据库容器的版本
export DB_TAG=v0.0.1
# API容器的版本
export API_TAG=v0.1.19
# 前端容器的版本
export DASHBOARD_TAG=v0.1.16
# 对接大脑的容器的版本
export REQ_TAG=v0.0.5
# 图像识别服务的版本
export RECOGNIZE_TAG=v0.1
# 回放服务的版本
export REPLAY_TAG=v0.0.1
# 镜像源的前缀,外网无需填写,内网按需填写
export PREFIX=

# 停止服务
docker stop recognize_service yToolsBox-req-custom yToolsBox-replay-custom yToolsBox-dashboard-custom yToolsBox-api-custom yToolsBox-db
# 删除服务
docker rm recognize_service yToolsBox-req-custom yToolsBox-replay-custom yToolsBox-dashboard-custom yToolsBox-api-custom yToolsBox-db
# 删除网络与挂载卷
docker network rm ytoolsbox_network
docker volume rm ytoolsbox_db-data


# 新建网络与挂载卷
docker network create --driver bridge ytoolsbox_network
docker volume create ytoolsbox_db-data
# 拉起服务
docker run -itd -p 5432:5432 --name yToolsBox-db --network ytoolsbox_network -e POSTGRES_PASSWORD=test123456 -v ytoolsbox_db-data:/var/lib/postgresql/data  ${PREFIX}yanqiaoyu/ytoolsbox-db:${DB_TAG}
docker run -itd -p 8081:8081 --name yToolsBox-api-custom --network ytoolsbox_network -e HOST_SCRIPT_PATH=/home/yToolsBox/api/Script -v /home/yToolsBox/api/Script:/root/Script ${PREFIX}yanqiaoyu/ytoolsbox-api-custom:${API_TAG} supervisord -c /etc/supervisord.conf
docker run -itd -p 80:80 --name yToolsBox-dashboard-custom --network ytoolsbox_network  ${PREFIX}yanqiaoyu/ytoolsbox-dashboard-custom:${DASHBOARD_TAG}
docker run -itd -p 2468:2468 --name yToolsBox-req-custom --network ytoolsbox_network  ${PREFIX}yanqiaoyu/ytoolsbox-req-custom:${REQ_TAG}
docker run -itd -p 3579:3579 --name recognize_service --network ytoolsbox_network  ${PREFIX}yanqiaoyu/recognize_service:${RECOGNIZE_TAG}
docker run -itd --name yToolsBox-replay-custom --net host ${PREFIX}yanqiaoyu/ytoolsbox-replay-custom:${REPLAY_TAG}