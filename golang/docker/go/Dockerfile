####################### Local stage #######################
FROM golang:1.19.1-alpine AS development

# gccのインストール
RUN apk add --no-cache gcc musl-dev git

# 作業ディレクトリの作成と設定
WORKDIR /app

# 依存関係のインストール
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# airのインストール
RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/99designs/gqlgen@latest && \
  go install github.com/volatiletech/sqlboiler/v4@latest && \
  go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest && \
  go install -tags mysql github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . .

CMD ["air"]

####################### Builder stage #######################
FROM development AS builder

# アプリケーションをビルド
RUN go build -trimpath -ldflags "-w -s" -o main /app/cmd/server.go

####################### Runtime stage #######################
FROM alpine:3.13

# 必要なパッケージをインストール
RUN apk add --no-cache ca-certificates

# ビルダーステージからバイナリをコピー
COPY --from=builder /app/main .

# # アプリケーションの実行
CMD ["./main"]
