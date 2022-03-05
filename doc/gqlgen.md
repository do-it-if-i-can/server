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

## define schema
gqlgenはスキーマファーストライブラリのため、コードを書く前にGraphQLスキーマ定義言語を使用しAPIを記述します。
デフォルトでは`schema.graphql`というファイルに記述します。
これは必要に応じて分割することが可能です。

## implement the resolvers
`go run github.com/99designs/gqlgen generate` gqlgenのgenerateコマンドを実行するとスキーマファイルがmodelと比較され
可能な場合はmodelに直接バインドされます。
自動生成直後のリゾルバは実装がないのでやっていく。(DBではなく一旦メモリ内への格納)

**resolver.go**
```go
package graph

import "github.com/do-it-if-i-can/server/graph/model"

type Resolver struct {
	todos []*model.Todo
}
```

`resolver.go`にアプリの依存関係を宣言する。

**schema.resolvers.go**
```go
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
```

## startup server
```
go run server.go
```

→createTodo sample mutation
```
mutation createTodo {
  createTodo(input: { text: "todo", userId: "1" }) {
    user {
      id
    }
    text
    done
  }
}
```

→list todos sample query
```
query findTodos {
  todos {
    text
    done
    user {
      name
    }
  }
}
```

# エラー解決集
## ResolverRootに変なのが入り込む場合
- goのmodelパッケージの中のstructで大文字から初めてないやつがないか確認する