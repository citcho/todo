# 最強のTodoアプリ

![go_version](https://img.shields.io/badge/go-1.22.0-blue.svg?logo=go) ![react_version](https://img.shields.io/badge/react-18.0.37-blue?logo=react) ![nodejs_version](https://img.shields.io/badge/node.js-18.16.0-blue.svg?logo=nodedotjs) ![npm_version](https://img.shields.io/badge/npm-9.5.1-blue.svg?logo=npm) ![nginx_version](https://img.shields.io/badge/nginx-1.25-blue.svg?logo=nginx) ![mysql_version](https://img.shields.io/badge/mysql-8.0-blue.svg?logo=mysql) ![terraform_version](https://img.shields.io/badge/terraform-1.7.5-blue.svg?logo=terraform)

最強に赤字になりそうなTodoアプリです
React, GoのSPAで作成しています

【機能】
- サインアップ
- サインイン
- サインアウト
- Todo作成
- Todo完了
- Todo一覧

【こだわったところ】
- ヘキサゴナルアーキテクチャ
- CQS
- 2月リリースのGo1.22で機能追加されたnet/httpパッケージを使用している
- echoやginといったライブラリは使用していない
- 隠れ人気のORMライブラリbunの採用
- JWTを使用した認証認可機能
- Atomic Design（ちょっとだけ）
- 無駄のないCSS笑
- Github Actionsを使用した自動化（単体テスト、テーブル定義書・ER図の自動作成、Reactの自動ビルド・デプロイ、Goイメージの自動ビルド・デプロイ）
- TerraformでのIaC化
- Goイメージのマルチステージビルド

【時間なくて断念したところ】
- CQRS（DynamoDBとRDSで分けてみたかった）
- 統合テスト
- bulletproof-react
- OpenAPI（Swagger）
- 静的解析ツールの導入、自動化

【感想】
楽しかった（KONAMI）

## 環境構築

### ①Docker
```bash
# コンテナ起動
make up
make gen
make gen-key

```

### ②DBマイグレーション
```bash
make migrate
```

### ③/etc/hosts
```
127.0.0.1 dev-todo.citcho.com
127.0.0.1 api.dev-todo.citcho.com
```

## テスト関連

### テスト実行コマンド
```bash
make test
```
### テストファイル雛形作成コマンド
```bash
# example ※テストファイルを作成したいディレクトリに移動する必要あり
$ gotests -w -all greeter.go
```

### モックファイル作成コマンド
```bash
# モックが必要なインターフェースには、以下のように定義の上にgenerate文を記述する
# //go:generate mockgen -source=./usecase.go -destination=./mock/usecase.go

# 書いたら下記実行
make gen
```

## デバッグ関連
基本的には`log`パッケージでデバッグします。
