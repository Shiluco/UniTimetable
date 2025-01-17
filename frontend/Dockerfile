FROM node:22.13.0-alpine

# 作業ディレクトリを作成
WORKDIR /app

# yarnを有効化
RUN corepack enable yarn

# package.json, yarn.lock, .yarnrc.yaml をコピー
COPY package.json yarn.lock .yarnrc.yaml ./

# Yarnのキャッシュディレクトリを作成
RUN mkdir -p .yarn && chmod -R 777 .yarn

# 開発用依存関係をインストール
RUN --mount=type=cache,target=/root/.cache/yarn/v6 \
    yarn install

# 開発サーバーのポートを公開
EXPOSE 3000

# 開発用サーバを起動
CMD ["yarn", "dev"]
