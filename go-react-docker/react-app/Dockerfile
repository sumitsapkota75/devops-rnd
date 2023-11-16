FROM node:20-alpine

RUN yarn add nodemon

WORKDIR /app

COPY package.json ./

RUN yarn install

COPY . .

EXPOSE 4000

CMD ["yarn", "start"]