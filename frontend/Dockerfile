# ベースイメージを指定（Node.jsのバージョンは11.x系に合わせる）
FROM node:22.13.0-alpine

# 作業ディレクトリを作成
WORKDIR /app

# package.json と package-lock.json をコピー
COPY package*.json ./

# npmのバージョンを指定（11.0.0に固定）
RUN npm install -g npm@11.0.0

# 依存関係をインストール
RUN npm install

# アプリケーションのコードをコピー
COPY . .

# Next.js のビルド
RUN npm run build

# アプリケーションのポートを公開
EXPOSE 3000

# アプリケーションを起動
CMD ["npm", "run", "dev"]