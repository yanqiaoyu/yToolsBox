cd ./ytoolsbox-vue
npm run build
docker rmi yanqiaoyu/ytoolsbox-dashboard:v0.1 yanqiaoyu/ytoolsbox-api:v0.1 -f
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:v0.1
docker push yanqiaoyu/ytoolsbox-api:v0.1
