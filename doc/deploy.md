# CloudRunへのデプロイ
ref: https://zenn.dev/ucwork/articles/eb4242ba124581#%E6%9C%AC%E7%95%AA%E3%83%87%E3%83%97%E3%83%AD%E3%82%A4

**※アプリケーションを変更して再度デプロイするとき(後でCI化予定)**
dockerコマンドたちを再実行して以下のコマンドを実行する。
(terraformはあくまで環境をcode化するもので、デプロイのたびにterraform applyするものではない)

```
gcloud run deploy cloudrun-srv --image=gcr.io/qin-todo-341312/qin-todo-api:latest --region=asia-northeast1
```

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

### unauthorized: You don't have the needed permissions to perform this operation, and you may have invalid credentials. To authenticate your request, follow the steps ...

がdocker pushで出たので記載のURLにそって権限をよしなにやる
https://cloud.google.com/container-registry/docs/advanced-authentication

`gcloud`コマンド使ってログインしたり`docker-credential-gcr`をinstallしたりする必要があるっぽい？
ref: https://genzouw.com/entry/2022/02/05/080015/2918/
ref: https://zenn.dev/phi/articles/gcloud-setup-with-homebrew-on-mac

**`gcloud`使えるようにする**

```
brew install google-cloud-sdk --cask

==> Linking Binary 'bq' to '/opt/homebrew/bin/bq'
==> Linking Binary 'docker-credential-gcloud' to '/opt/homebrew/bin/docker-credential-gcloud'
==> Linking Binary 'gcloud' to '/opt/homebrew/bin/gcloud'
==> Linking Binary 'git-credential-gcloud.sh' to '/opt/homebrew/bin/git-credential-gcloud'
==> Linking Binary 'gsutil' to '/opt/homebrew/bin/gsutil'
==> Linking Binary 'anthoscli' to '/opt/homebrew/bin/anthoscli'
🍺  google-cloud-sdk was successfully installed!
```

```
❯ gcloud --version
Google Cloud SDK 372.0.0
```

**googleアカウントと連携する**

```
gcloud auth login

==>You are now logged in as [s.kawabe2281@gmail.com].
==>Your current project is [None].  You can change this setting by running:
==> $ gcloud config set project PROJECT_ID
```

**GCPプロジェクトをセットする**

```
gcloud config set project qin-todo-341312
```

**GCR に対してビルドイメージをpushできるよう認証を行う**

```
gcloud auth configure-docker
```

` ~/.docker/config.json`に認証情報が保存された。
ここでdocker pushすると今度は成功したみたい🙆‍♂️


## terraformインストールする
https://learn.hashicorp.com/tutorials/terraform/install-cli

```
brew tap hashicorp/tap
```

```
brew install hashicorp/tap/terraform
```

```
❯ terraform -version
Terraform v1.1.5
on darwin_arm64
```

## cloud runにデプロイする

```
cd environments/production

terraform init # Terraform設定ファイルを標準形式とスタイルに書き換えるために使用。
terraform validate # terraformファイルの構文を検証
==> Success! The configuration is valid.
terraform plan # 実行計画の作成
terraform apply # planの計画に沿って動作を適用する
```

### planで怒られる
providerという記述が足りてないようなので追加した

```
provider "google" {
  project     = "qin-todo-341312"
  region      = "asia-northeast1"
}
```

上記でも怒られるので以下を実行する必要があるらしい

```
gcloud auth application-default login
```

> こちらはGoやら各言語でのSDKを使ったプログラムを実行する際の認証を得るために使います。
> https://christina04.hatenablog.com/entry/gcp-auth

認証情報がローカルに保存されたのでこうすればいける？
→terraform plan通ったっぽい

### applyで怒られる
→Cloud Run自体を有効化していなかった