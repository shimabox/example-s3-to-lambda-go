# example-s3-to-lambda-go

## ユースケース

挨拶されたら挨拶された内容を保存する。

### 技術的ユースケース

挨拶(S3へjsonファイルをPut)されたら挨拶された内容(jsonファイルの中身)を保存(標準出力してCloudWatchへ保存)する。

- 挨拶
  - S3へjsonファイルをPut
- 挨拶された内容
  - jsonファイルの中身
- 保存
  - 標準出力してCloudWatchに残す

## リソース

※ 名前などは環境変数(`.env`)で定義

### LocalStack

- ローカルでのAWS開発用
- [Installation | Docs](https://docs.localstack.cloud/getting-started/installation/#docker-compose "Installation | Docs")

### Lambda

- 関数
  - example-s3-to-lambda-go 

### S3

- Lambda関数配置バケット
  - `local-lambda-storage-bucket`
- Putイベント監視バケット
    - `local-input-bucket`

### CloudWatch

- LogGroup
  - `/aws/lambda/example-s3-to-lambda-go` 

## Usage

### clone

```shell
git clone https://github.com/shimabox/example-s3-to-lambda-go.git
cd example-s3-to-lambda-go
cp .env.example .env
```

### build

```shell
make build
make cl # create lambda.
```

### test

```shell
# all test
make alltest

# feature test
make featuretest

# unit test
docker compose exec go make gotest
// or locally.
make gotest
```

### Lint

```shell
docker compose exec go make lint
// or locally.
make lint
```

## Command

- Container up
  - `make up`
- Container down
  - `make down`
- Container restart
  - `make restart`
- Update lambda.
  - `make ul # update lambda`
- Check lambda operation.
  - `docker compose exec localstack bash scripts/operation_check.sh`

## Todo

- [x] とりあえず動くように
- [x] xxxアーキテクチャでテストを書きやすくする
  - オレオレアーキテクチャ
- [x] localstackでの動作確認
- [ ] SAMでデプロイ
- [ ] GitHub Actions
