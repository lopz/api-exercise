#!/bin/bash

set +e

version=$1

IMAGE_NAME="lowlifebob/csapi"
TAG="alpine-1.0.1"

echo "Building image ${IMAGE_NAME}:${TAG}"

docker build -t csapi-build-base:alpine-1.0.1 --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` -f deployments/docker/Dockerfiles/build-base/Dockerfile .
docker build -t csapi-build-csapi:alpine-1.0.1 --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` -f deployments/docker/Dockerfiles/build-csapi/Dockerfile .
docker build -t ${IMAGE_NAME}:${TAG} --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` -f deployments/docker/Dockerfiles/csapi-mongo/Dockerfile .
docker build -t ${IMAGE_NAME}:seeder-${TAG} --build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` -f deployments/docker/Dockerfiles/mongo-seed/Dockerfile .

echo "Pushing images to dockerhub."
docker push ${IMAGE_NAME}:${TAG}
docker push ${IMAGE_NAME}:seeder-${TAG}