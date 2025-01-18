# ベースイメージを指定（Node.jsのバージョンは11.x系に合わせる）
FROM node:22.13.0-alpine

# 作業ディレクトリを作成
WORKDIR /app

# package.json と yarn.lock をコピー
COPY package.json yarn.lock ./


# 依存関係をインストール
RUN yarn install

# Storybook の初期化（必要な場合）
RUN npx sb init

# アプリケーションのコードをコピー
COPY . .

# Next.js のビルド
RUN yarn build

# Storybook のビルド
RUN yarn build-storybook

# アプリケーションと Storybook のポートを公開
EXPOSE 3000 6006

# デフォルトで Next.js を起動
CMD ["yarn", "dev"]

# Storybook を個別で起動する場合の例
# docker run -p 6006:6006 イメージ名 yarn storybook
