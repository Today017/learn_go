# learn_go

[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)

2025/12/13勉強開始 Thanks to そま

## 1
### 1-3
#### HTTPメソッド
- GET データを取得
- POST データをクライアントからサーバーへ
- PUT データの書き換え
- DELETE データの削除
(あくまで慣例ではある)
現時点ではどんなHTTPメソッドも受け付けるようになっているので、特定のメソッドだけ受け付けるように変えてみる

#### HTTPレスポンスステータスコード
- 200 OK
- 400 Bad Request ユーザーのリクエストの値が不正
- 403 Forbidden
- 404 Not Found
- 405 Method Not Allowed 
- 500 Internal Server Error

```bash
curl http://localhost:8080/hello -X DELETE -w '%{http_code}\n'
```
- `-w '%{http_code}\n'` オプションでレスポンスステータスコードを表示できる

### 1-4 gojilla/muxパッケージ
- ルーティングを簡単に実装できるパッケージ

```bash
go get -u github.com/gorilla/mux
```

`-u` で最新バージョンを取得できる

- `go.mod`
    - モジュールのルートディレクトリを示す
    - パッケージの依存関係を記録する
- `go.sum`
    - 依存パッケージのバージョンとチェックサムを記録する
    - チェックサム：ビルドの整合性を確認するために使われる
        - モジュールをインストールする際、Googleが管理しているチェックサムと照合して改ざんされていないか確認する
        - https://go.dev/ref/mod#authenticating

### 1-5 パスパラメータ
- URLに変数を含めてハンドラ関数に引数として渡す仕組み
- 正規表現を使う
  - `mux.Vars(req)` でパスパラメータを取得する
- `strconv` パッケージ
  - 文字列と他の基本データ型（int、floatなど）を相互変換するためのパッケージ
  - `strconv.Atoi` で文字列をintに変換できる

### 1-6 クエリパラメータ
- URLの?以降に含まれるキーと値のペア
    - `https://kenkoooo.com/atcoder/#/user/Today03?userPageTab=submissions`
    - の `?userPageTab=submissions` の部分みたいなやつ
- ユーザー側でkey-valueのセットを自由にURLに追加できる

```go
type Request struct {
    URL *url.URL
    ...
}

type URL struct {
    ...
    Scheme  string // http, https
    Host    string // example.com
    Path    string // /atcoder/#/user/Today03
    RawQuery string // userPageTab=submissions
    Fragment string // #以下の文字列
}
```

- `func (u *URL) Query() Values`
    - URL構造体のメソッド
    - map[string][]string型を返す


### 1章まとめ
- 基本的なWebサーバーの実装方法を勉強した
- WebサーバーがHTTPリクエストを受け取ってレスポンスを返すまでの流れを理解した
    - ルーター：HTTPリクエストのURLパスに基づいてハンドラ関数を呼び出す
    - ハンドラ関数：HTTPリクエストを処理してレスポンスを返す
    - HTTPメソッドの制限
    - パスパラメータとクエリパラメータの取得方法

## 2
### 2-1 構造体
- HTTPリクエストやレスポンスのデータを構造体で定義する

`time.Time` 型
- 日付と時刻を表すGoの標準ライブラリの型

構造体の出力
```go
fmt.Printf("%+v\n", article)
```
- `%+v` で構造体のフィールド名と値を表示できる

### 2-2 json
`func Marshal(v any) ([]byte, error)` Goの構造体からjsonを作るための関数

