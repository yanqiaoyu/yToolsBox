# 获取编译时间（格式：YYYYMMDDHHMMSS）
BUILD_TIME=$(date +"%Y%m%d%H%M%S")

# 数据库容器的版本
export DB_TAG=v0.0.1
# API容器的版本
export API_TAG="v3.0.21-${BUILD_TIME}"
# 前端容器的版本
export DASHBOARD_TAG="v3.0.21-${BUILD_TIME}"
# 对接大脑的容器的版本
export REQ_TAG=v3.0.18
# 图像识别服务的版本
export RECOGNIZE_TAG=v0.1
# 回放服务的版本
export REPLAY_TAG=v0.0.1

cd ./ytoolsbox-vue-custom
npm run build
cd ../
# 只编译,不运行服务
# docker-compose up -d
docker-compose build
docker push yanqiaoyu/ytoolsbox-api-custom:${API_TAG}
docker push yanqiaoyu/ytoolsbox-dashboard-custom:${DASHBOARD_TAG}
docker push yanqiaoyu/ytoolsbox-req-custom:${REQ_TAG}
