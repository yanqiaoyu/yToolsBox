docker stop yToolsBox-api yToolsBox-dashboard yToolsBox-db
docker rm yToolsBox-api yToolsBox-dashboard yToolsBox-db

docker images | grep none | awk '{print $3}' | xargs docker rmi
cd ./ytoolsbox-vue
npm run build
docker rmi yanqiaoyu/ytoolsbox-dashboard:v0.2 yanqiaoyu/ytoolsbox-api:v0.2 -f
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:v0.2
docker push yanqiaoyu/ytoolsbox-api:v0.2
