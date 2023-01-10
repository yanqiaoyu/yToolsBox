export TAG=v0.0.1
docker build -t yanqiaoyu/ytoolsbox-db-custom:${TAG} .
docker push yanqiaoyu/ytoolsbox-db-custom:${TAG}
