export TAG=v0.2.2

docker-compose down -v
cd ./ytoolsbox-vue
npm run build
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:${TAG}
docker push yanqiaoyu/ytoolsbox-api:${TAG}
