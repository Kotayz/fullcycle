#!make

start-docker:
	sudo service docker start

remove-all:
	docker rm -f $(shell docker ps -a -q)

remove-all-images:
	docker rmi -f $(shell docker images -a -q)