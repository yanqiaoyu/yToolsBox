export TAG=v0.2.1

docker-compose down all -v
cd ./ytoolsbox-vue
npm run build
cd ../
docker-compose up -d

docker push yanqiaoyu/ytoolsbox-dashboard:${TAG}
docker push yanqiaoyu/ytoolsbox-api:${TAG}
