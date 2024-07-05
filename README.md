# TODO SPA

![go_version](https://img.shields.io/badge/go-1.22.0-blue.svg?logo=go) ![react_version](https://img.shields.io/badge/react-18.0.37-blue?logo=react) ![nodejs_version](https://img.shields.io/badge/node.js-18.16.0-blue.svg?logo=nodedotjs) ![npm_version](https://img.shields.io/badge/npm-9.5.1-blue.svg?logo=npm) ![nginx_version](https://img.shields.io/badge/nginx-1.25-blue.svg?logo=nginx) ![mysql_version](https://img.shields.io/badge/mysql-8.0-blue.svg?logo=mysql) ![terraform_version](https://img.shields.io/badge/terraform-1.7.5-blue.svg?logo=terraform)

## Features
### Auth
- サインアップ
- サインイン
- サインアウト

### TODO
- Todo作成
- Todo完了
- Todo一覧

## Architecture
```bash
.
├── Makefile  # タスクランナー
├── README.md
├── _tools    # 設定ファイルなど
│   ├── mysql
│   └── nginx
├── cmd       # バックエンドのエントリポイント
│   └── todo
├── compose.yml
├── deploy    # インフラ系
├── docs      # ドキュメント
├── internal  # cmdで使用する内部ライブラリ
│   ├── pkg      # 内部ライブラリ共有モジュール
│   │   ├── auth
│   │   ├── clock
│   │   ├── config
│   │   ├── database
│   │   ├── encrypt
│   │   ├── server
│   │   └── ulid
│   ├── todo
│   │   ├── adapter      # インフラ層
│   │   ├── app          # アプリケーション層
│   │   ├── domain       # ドメイン層
│   │   └── presentation # プレゼンテーション層
│   └── user
│       ├── adapter
│       ├── app
│       ├── domain
│       └── presentation
└── web # フロントエンド
    ├── env
    ├── favicon.ico
    ├── index.html
    ├── jsconfig.json
    ├── node_modules
    ├── package-lock.json
    ├── package.json
    ├── src
    │   ├── components # コンポーネント管理
    │   ├── hooks      # 汎用的なカスタムフック
    │   ├── libs       # node_modulesのラッパー関数
    │   ├── main.jsx
    │   ├── routes     # ルーティング設定
    │   ├── schemas    # バリデーションスキーマ
    │   ├── stores     # global state
    │   └── styles     # 汎用的なスタイル
    └── vite.config.js
```

- クリーンアーキテクチャ
- CQS
- Atomic Design
- Github Actions
  - テーブル定義書・ER図
  - 単体テスト
  - Reactのビルド/デプロイ
  - APIイメージのビルド/デプロイ

## 環境構築
### バックエンド
```bash
make up
make gen-key
```

### フロントエンド
```bash
cd web
npm run dev
```

### マイグレーション
```bash
make migrate
```

### ドメイン設定
```bash
# /etc/hosts
127.0.0.1 dev-todo.citcho.com
127.0.0.1 api.dev-todo.citcho.com
```

## テスト
```bash
# テスト実行
make gen
make test

# テストファイル雛形作成
gotests -w -all greeter.go
```

## デバッグ
```bash
make logs
```

```go
foo := &Foo{Bar: "baz"}
log.Printf("foo: %s", foo.Bar)
```
