#!/bin/sh

IMAGE=mongo:3.6.8
NAME=mongo_dev

docker run -ti --rm \
	--name $NAME \
	--net host \
	--env MONGO_INITDB_ROOT_USERNAME=admin \
	--env MONGO_INITDB_ROOT_PASSWORD=123456 \
	--env DB_USERNAME=app \
	--env DB_PASSWORD=654321 \
	--env AUTHDB=admin \
	--env DB_NAME=products \
	--env COLLECTION=docs \
	--publish 27017:27017 \
	--volume $PWD/data:/data\
	--volume $PWD/init:/docker-entrypoint-initdb.d \
	$IMAGE
