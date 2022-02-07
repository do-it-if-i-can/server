# GraphQLのスキーマとは
GraphQL APIの仕様を表現するもの。
スキーマ定義言語(SDL)を用いて表現する。

**スキーマ定義の例**

```
type Book {
  is: ID!
  name: String
  pageCount: Int
  auther: Auther
}

type Auther {
  id: ID!
  firstName: String
  lastName: String
}
```

---

# GraphQLの型

## スカラー型
基本的な型

- String
- Int
- Float
- Boolean
- ID

## オブジェクト型
1つ以上のスキーマで定義されているフィールドの集合をオブジェクト型という。
上記で表した`auther: Auther`のような型がそれにあたる。

## 列挙型
色んな言語のenumとほぼ同じ

```
enum Color {
    RED
    YELLOW
    BLUE
    GREEN
}

type Apple {
    color: Color
}
```

## リスト
フィールドには型のリストを定義することも可能。
以下はStringのリストを持つスキーマの例。

```
type Department {
  id: ID!
  memberName: [String]
}
```

## ルート型 🤩
GraphQLのクエリはルート型で始まる。
ルート型とはデータソースに対する操作を表現する型。
- Query
- Mutation
- Subscription

## ユニオン型
Union型は複数の型のうち一つを返す型。

```
union AgendaItem = StudyGroup | Workout

type StudyGroup {
    name: String!
    subject: String
    students: [User!]
}

type Workout {
    name: String!
    reps: Int!
}

type Query {
    agenda: AgendaItem
}
```

---

# queryの引数の指定
https://graphql.org/learn/schema/#arguments

```
type Starship {
  id: ID!
  name: String!
  length(unit: LengthUnit = METER): Float
}
```

全てのフィールドは0個以上の引数を含めることができる。
全ての引数には名前がつけられている。
JSやPythonなどと異なり、GraphQLの引数は具体的に名前で渡される。

また、上の例のようにデフォルト引数を受け取ることも可能。

# エクスクラメーションマークについて
！マークがついているフィールドは、「そのフィールドがnullになることがない」ことを表す。

## リストのエクスクラメーションマーク
- [String]
  - nullかもしれないリストで中身もnullかもしれない文字列型
- [String!]
  - nullかもしれないリストで中身はnullではない文字列型
- [String]!
  - nullではないリストで中身はnullかもしれない文字列型
- [String!]!
  - nullではないリストで中身もnullではない文字列型