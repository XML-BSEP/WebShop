FROM node:10-alpine as build-step

RUN mkdir -p /app

WORKDIR /app

COPY package.json /app
RUN npm install
RUN npm prune 

RUN npm install -g @angular/cli@7.3.9

COPY . /app

EXPOSE 4300

CMD ng serve --o --host 0.0.0.0
