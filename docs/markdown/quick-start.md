# Quick Start 

1. dockerでgoを管理する場合
2. localでgoを管理する場合

# 共通初期設定
## env file作成
localかdockerはそれぞれの環境に合わせて
```shell
$ cp ./cmd/app/.env.<local|localdocker> ./cmd/app/.env
$ cp ./cmd/cli/.env.<local|localdocker> ./cmd/cli/.env
```

# 1. dockerバージョン

## コンテナ実行
```shell
$ make oapi/gen/app
$ make docker/up
# コンテナ内実行コマンド(docker desktop内で実行出来れば何でも良い)
$ make docker/exec/go
$ make go/install/tools
$ make ent/gen
$ make oapi/codegen/app
$ make mockgen
$ go run ./cmd/cli/main.go -exec genJwtSecretKeyForApp
$ go run ./cmd/cli/db/main.go -query migrate
$ go run ./cmd/cli/db/main.go -query seed
$ air -c ./cmd/app/air.toml
```

## アクセスURL
サインアップのエンドポイント  
http://localhost:8080/signup

# 2. localバージョン

## コンテナ起動
```shell
$ make docker/up/db
$ make docker/up/mailhog
```

## ツールのインストールとAirサーバー起動
```shell
$ go mod tidy
$ make go/install/tools
$ make ent/gen
$ make oapi/gen/app && make oapi/codegen/app
$ make mockgen
$ go run ./cmd/cli/main.go -exec genJwtSecretKeyForApp
$ go run ./cmd/cli/db/main.go -query migrate
$ go run ./cmd/cli/db/main.go -query seed
$ air -c ./cmd/app/air.toml
```
※ airが利用できない場合は、ターミナル再起動

## アクセスURL
サインアップのエンドポイント  
http://localhost:8080/signup


