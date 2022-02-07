# おぼえがき

# gqlgen getting started
https://gqlgen.com/getting-started/

## project setup

```
mkdir gqlgen-todos
cd gqlgen-todos
go mod init github.com/[username]/gqlgen-todos
```

→ tools.goに以下をimport
```go
package tools

import (
	_ "github.com/99designs/gqlgen"
)
```

```
go mod tidy
```

## building the server

```
go run github.com/99designs/gqlgen init
```

スケルトンを作成、以下はデフォルトの自動生成系ディレクトリ構成。
`gqlgen.yml`を編集すると色々設定できる


```
├── go.mod
├── go.sum
├── gqlgen.yml               - gqlgenの設定ファイル、生成されるコードを制御する
├── graph
│   ├── generated            - ランタイムによって生成されるもののみを含むパッケージ
│   │   └── generated.go
│   ├── model                - 自動生成&手動生成の全てのgraphql model用パッケージ?
│   │   └── models_gen.go
│   ├── resolver.go          - 自動では再生成されないファイル、必要な依存性注入を後から書き足す
│   ├── schema.graphql      - スキーマを記述する。スキーマは複数のgraphqlファイルに分割することもできる
│   └── schema.resolvers.go  - schema.graphqlのリゾルバ実装
└── server.go                - アプリケーションのエントリポイント。自由にカスタマイズ可能
```

**※ GraphQLにおけるResolverとは**
スキーマが単なるmodelの定義を表す用語なのに対し、リゾルバは実際のデータ操作を行うものを指す。
リゾルバの実態は特定のフィールドのデータを返す関数。

