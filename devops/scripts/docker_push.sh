#!/bin/bash
docker build -t zinefer/habits .
echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push zinefer/habits