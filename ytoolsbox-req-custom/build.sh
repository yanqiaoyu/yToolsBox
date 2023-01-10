export TAG=v0.0.2
docker build -t yanqiaoyu/ytoolsbox-req-custom:${TAG} .
docker push yanqiaoyu/ytoolsbox-req-custom:${TAG}
