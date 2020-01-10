#!/bin/bash
commit=$(git rev-parse HEAD | cut -c 1-7)
docker build --build-arg COMMIT=$commit -t zinefer/habits .
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push zinefer/habits