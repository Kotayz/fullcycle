FROM node:18.14.2-alpine3.17

WORKDIR /usr/src/app

COPY . .

RUN npm install

EXPOSE 3000

ENTRYPOINT [ "node", "index.js" ]