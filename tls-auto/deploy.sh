#!/bin/bash -ex

TAG=v0.6.0
docker build -t server:${TAG} -f server/Dockerfile .
docker build -t client:${TAG} -f client/Dockerfile .

#k3d image import --cluster flux server:${TAG}
#k3d image import --cluster flux client:${TAG}

# Update yaml
sed "s,IMAGE,server\:$TAG," "server/server.yaml" > config/server.yaml
sed "s,IMAGE,client\:$TAG," "client/client.yaml" > config/client.yaml

kubectl apply -f config/
