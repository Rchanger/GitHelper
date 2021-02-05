FROM node:12.2.0-alpine

COPY . /home/GitHelper/client/
WORKDIR /home/GitHelper/client/

RUN npm install

RUN npm install -g @vue/cli

EXPOSE 3030

CMD ["npm", "run", "dev"]