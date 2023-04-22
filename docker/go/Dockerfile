####################### Local stage #######################
FROM golang:1.19.1-alpine AS development

# 環境変数の設定
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# gccのインストール
RUN apk add --no-cache gcc musl-dev git

# 作業ディレクトリの作成と設定
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR /app

# 依存関係のインストール
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# airのインストール
RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air"]

####################### Builder stage #######################
FROM development AS builder

# アプリケーションをビルド
RUN go build -trimpath -ldflags "-w -s" -o main /app/server.go

####################### Runtime stage #######################
FROM alpine:3.13

# 必要なパッケージをインストール
RUN apk add --no-cache ca-certificates

# 開発ステージからreflexバイナリをコピー
COPY --from=development /go/bin/ /usr/local/bin/

# ビルダーステージからバイナリをコピー
COPY --from=builder /app/main .

# # アプリケーションの実行
CMD ["./main"]
