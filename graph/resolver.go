package graph

import "github.com/do-it-if-i-can/server/graph/model"

// This file will not be regenerated automatically.
// このファイルは自動的に再生成されません

// It serves as dependency injection for your app, add any dependencies you require here.
// アプリの依存性注入の役割を果たすもので、必要な依存性をここに追加します。

type Resolver struct {
	todos []*model.Todo
}
