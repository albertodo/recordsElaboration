FROM node:10-alpine AS vue-cli
RUN npm install --loglevel=error -g @vue/cli

FROM node:10-alpine AS npm-install
USER  node 
WORKDIR /home/node
COPY package.json package.json
COPY package-lock.json package-lock.json
RUN npm ci --loglevel=error

FROM node:10-alpine AS vue-build
COPY --from=vue-cli /usr/local/lib/node-modules/ /usr/local/lib/node-modules/
USER node
WORKDIR /home/node
COPY --from=npm-install /home/node .
COPY vue.config.js vue.config.js
COPY public public
COPY src src
RUN npm run build

FROM node:10-alpine
RUN npm install --loglevel=error -g serve
USER node
WORKDIR /home/node
COPY --from=vue-build /home/node/dist /home/node/dist
EXPOSE 8080
CMD ["serve","-s","dist","-p","8080"]


