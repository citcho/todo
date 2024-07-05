# バイナリ作成用コンテナステージ
FROM golang:1.22.0-bookworm as deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/todo
RUN GOARCH=amd64 GOOS=linux go build -trimpath -ldflags "-w -s" -o ../../main

# ------------------------------------------------------------

# デプロイ用コンテナ
FROM debian:bookworm-slim as deploy

RUN apt-get update

WORKDIR /app

COPY --from=deploy-builder /app/main /app/

CMD ["./main"]

# ------------------------------------------------------------

# ローカル用ライブリロード対応コンテナステージ
FROM golang:1.22.0-bookworm as dev

WORKDIR /app

RUN go install -v golang.org/x/tools/gopls@latest \
    && go install -v github.com/rogpeppe/godef@latest \
    && go install github.com/golang/mock/mockgen@v1.6.0 \
    && go install github.com/air-verse/air@latest \
    && go install github.com/sqldef/sqldef/cmd/mysqldef@latest \
    && go install github.com/cweill/gotests/gotests@latest

CMD ["air", "-c", "./.air.toml"]
