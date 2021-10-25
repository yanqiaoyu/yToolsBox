docker-compose down --rmi all -v
cd ./ytoolsbox-vue
npm run build
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:v0.2
docker push yanqiaoyu/ytoolsbox-api:v0.2
