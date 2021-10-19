cd ./ytoolsbox-vue
npm run build
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:v0.1
docker push yanqiaoyu/ytoolsbox-api:v0.1