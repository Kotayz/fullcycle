#!make

remove-all:
	docker rm -f $(shell docker ps -a -q)

remove-all-images:
	docker rmi -f $(shell docker images -a -q)