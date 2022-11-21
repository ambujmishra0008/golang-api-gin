#!/bin/bash

api_name='track_order_detail'
api_version=`date +%Y%m%d%H%M%S`

docker build -t localhost:6666/$api_name:"Release_"$api_version .
docker push localhost:6666/$api_name:"Release_"$api_version

sed -i 's/Release_[0-9]\{1,\}/Release_'"$api_version"'/' deployment_DEV.yml
kubectl apply -f deployment_DEV.yml

docker system prune -f
