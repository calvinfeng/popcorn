FROM node:10
LABEL authors="Calvin Feng"

COPY . /api

WORKDIR /api/ui
RUN npm install
RUN npm run build
RUN npm run min

WORKDIR /api/server
RUN npm install

ENV BACKEND_HOST=backend
EXPOSE 8000

CMD node app.js
