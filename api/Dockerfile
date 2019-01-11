FROM node:10
LABEL authors="Calvin Feng"

COPY . /api
WORKDIR /api

EXPOSE 8000

RUN npm install

CMD node index.js
