# CloudRunへのデプロイ
ref: https://zenn.dev/ucwork/articles/eb4242ba124581#%E6%9C%AC%E7%95%AA%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4

## GCRでプロジェクト作成しGCP有効にする

1. qin-todoでプロジェクト作成
2. Google Container Registry API コンソール画面に行き「有効にする」選択

### Container Registryのタグの構成

```
gcr.io/プロジェクトID/コンテナイメージの名前
```

## Docker tag作成しGCRへpush

ref: https://www.topgate.co.jp/gcp08-how-to-use-docker-image-container-registry

```
# --targetでビルドステージを指定し、-tでタグをつける
docker build --platform linux/amd64 --target prod -t qin-todo-api .

# 「qin-todo-api」というイメージIDでイメージにタグ付けし、URLを指定(:○○の部分はバージョンのラベル)
docker tag qin-todo-api gcr.io/qin-todo-341312/qin-todo-api:latest

# docker pushでGCRにイメージをpushする
docker push gcr.io/qin-todo-341312/qin-todo-api:latest
```

- 「qin-todo-api」がdocker buildしたあとのIMAGE IDとなるもの
- 「gcr.io/qin-todo-341312/qin-todo-api:latest」がレジストリのURL
- 基本レジストリにpushする際は上のような流れになる気がする「docker build」→「docker tag」→「docker push」

