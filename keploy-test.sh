#!/bin/bash

set -e

services=("user-svc" "order-svc" "payment-svc" "notification-svc")

declare -A ports
ports=( ["user-svc"]=8001 ["order-svc"]=8002 ["payment-svc"]=8003 ["notification-svc"]=8004 )

echo "Building Go services"
for svc in "${services[@]}"; do
  echo "Building $svc..."
  cd $svc
  go build -o app
  cd ..
done

echo "Building Docker images..."
for svc in "${services[@]}"; do
  echo "Building $svc Docker image..."
  docker build -t $svc-img ./$svc
done

echo "Starting Keploy recording..."
for svc in "${services[@]}"; do
  port=${ports[$svc]}
  echo "Recording traffic for $svc on port $port..."
  keploy record -c "docker run -p $port:$port --name $svc-container --rm $svc-img" --buildDelay 10
  echo "$svc recorded."
done

echo "Running Keploy tests..."
for svc in "${services[@]}"; do
  port=${ports[$svc]}
  echo "Testing $svc on port $port..."
  keploy test -c "docker run -p $port:$port --name $svc-test-container --rm $svc-img" --delay 10 --buildDelay 10
  echo "$svc tested."
done

# ./keploy-test.sh