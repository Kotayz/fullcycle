#!make

build-node-full-cycle-rocks:
	docker build -t kotayz/node-full-cycle:latest ./node/full-cycle-rocks

build-nginx-full-cycle-rocks:
	docker build -t kotayz/nginx-full-cycle:latest ./nginx/full-cycle-rocks

remove-all:
	docker rm -f $(shell docker ps -a -q)

remove-all-images:
	docker rmi -f $(shell docker images -a -q)