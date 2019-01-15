FROM node:10
LABEL authors="Calvin Feng"

COPY . /api

WORKDIR /api/ui
RUN npm install
RUN npm run build
RUN npm run min

WORKDIR /api/server
RUN npm install

ENV PORT=8080

ENV GRPC_HOSTNAME=backend
ENV GPC_GRPC_HOSTNAME=something

EXPOSE 8080

CMD node app.js
