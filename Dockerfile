# ビルドステージ
FROM amd64/golang:1.22.0-alpine AS builder

# クロスコンパイルの設定
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64

# 必要なパッケージのインストール
RUN apk add --no-cache gcc musl-dev sqlite-dev

# 作業ディレクトリを設定
WORKDIR /app

# Goモジュールファイルをコピー
COPY go.mod ./
COPY go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN go build -o book-manager

# 実行ステージ
FROM alpine:3.18

# SQLiteを実行するために必要なライブラリをインストール
RUN apk add --no-cache sqlite-libs

# 作業ディレクトリを設定
WORKDIR /app

# ビルドしたバイナリをコピー
COPY --from=builder /app/book-manager /book-manager

# 静的ファイルのディレクトリを指定
VOLUME ["/app/static"]

# ポート番号を指定
EXPOSE 8080

# アプリケーションを実行
CMD ["/book-manager"]

