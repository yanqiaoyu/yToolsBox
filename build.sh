# 数据库容器的版本
export DB_TAG=v0.0.1
# API容器的版本
export API_TAG=v0.1.17
# 前端容器的版本
export DASHBOARD_TAG=v0.1.16
# 对接大脑的容器的版本
export REQ_TAG=v0.0.5
# 图像识别服务的版本
export RECOGNIZE_TAG=v0.1
# 回放服务的版本
export REPLAY_TAG=v0.0.1

cd ./ytoolsbox-vue-custom
npm run build
cd ../
docker-compose up -d
docker push yanqiaoyu/ytoolsbox-api-custom:${API_TAG}
docker push yanqiaoyu/ytoolsbox-dashboard-custom:${DASHBOARD_TAG}
