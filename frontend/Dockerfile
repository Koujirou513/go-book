FROM node:20-bullseye AS builder

WORKDIR /app

COPY package.json /app/
COPY package-lock.json /app/

RUN npm install npm@latest && npm install

COPY . /app/

# 開発用サーバー起動
CMD [ "npm", "run", "dev" ]

# 本番環境
FROM nginx:alpine

# ビルドしたReactアプリをNginxにコピー
COPY --from=builder /app/dist /usr/share/nginx/html 

EXPOSE 80

# Nginxをフォアグラウンドで起動するための設定
CMD [ "nginx", "-g", "daemon off;" ]