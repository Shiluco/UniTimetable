FROM golang:1.21-alpine

WORKDIR /app

# 必要なパッケージのインストール
RUN apk add --no-cache \
    gcc \
    musl-dev \
    git

# 依存関係のコピーとインストール
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# テスト実行用のコマンド
CMD ["go", "test", "-v", "./..."] 