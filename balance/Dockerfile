FROM node:18.16.1-bullseye

WORKDIR /app

RUN npm i -g @nestjs/cli

COPY package-lock.json package.json ./

RUN npm i

COPY . .

RUN npm run build

CMD ["npm", "start"]