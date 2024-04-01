# Goイメージをベースとする
FROM golang:1.22-alpine

# SQLiteを立ち上げるのに必要
ENV CGO_ENABLED=1

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
RUN go build -o /book-manager

# 静的ファイルのディレクトリを指定（あなたのアプリケーションに合わせて変更してください）
VOLUME ["/app/static"]

# アプリケーションを実行
CMD ["/book-manager"]
