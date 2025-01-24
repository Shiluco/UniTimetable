FROM golang:latest AS builder

WORKDIR /app

# 必要なツールのインストール
RUN go install entgo.io/ent/cmd/ent@latest

# 依存関係のコピーとダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードのコピー
COPY . .

# entの生成
RUN go generate ./ent

EXPOSE 8080

CMD ["go", "run", "cmd/server/main.go"]
