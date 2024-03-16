#!/usr/bin/env bash

docker compose -f docker-compose.kind.yaml build
docker compose -f docker-compose.kind.yaml push

kubectl delete -k ./kube
kubectl apply -k ./kube
