IMAGE=ttlzx/golang_1_11_builder:0.1
docker build . -t $IMAGE
docker push $IMAGE
