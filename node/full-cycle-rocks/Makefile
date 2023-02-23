#!make

node:
	docker run --rm -it --name node -v $(shell pwd)/:/usr/src/app -p 3000:3000 node bash

nginx:
	docker run --rm -it --name nginx -v $(shell pwd)/:/usr/src/app -p 8080:80 nginx bash