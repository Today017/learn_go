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

### 2-3
- Go構造体とjsonキーの命名規則
    - Go構造体はキャメルケース、jsonキーはスネークケース

- `json` タグを使ってGo構造体のフィールドとjsonキーの対応を指定できる

- これ関数のオプションとかで一括でできないのかな

### 2-4 HTTPレスポンスにjsonを書き込む
- データモデルを定義する専用のパッケージ(`model` パッケージ)を作成する

- レスポンス書き込み
  - `io.WriteString` ではなく `w.Write` メソッドを使う
  - `string` 型ではなく `[]byte` 型を受け取る
  - `string` と `[]byte` って何が違うんだろう
    - `string` は不変（immutable）で、`[]byte` は可変（mutable）
    - `string` はテキストデータ、`[]byte` はバイナリデータを扱うのに適している
    - `string` はUTF-8エンコードされたテキストを表すのに使われ、`[]byte` は任意のバイナリデータを表すのに使われる
    - `string` は文字列操作に便利なメソッドが豊富に用意されているが、`[]byte` はバイトレベルでの操作が必要な場合に使われる
    - `string` から `[]byte` への変換は `[]byte(s)`、`[]byte` から `string` への変換は `string(b)`

- `curl` でクエリパラメータを指定するときはURLを`""`で囲むこと

### 2-5 HTTPリクエストボディからjsonを読み込む
- HTTPリクエストの中身
    - ヘッダー : メタデータ
    - ボディ : リクエストの本体
    - メソッド : リクエストの種類（GET、POSTなど）
    - URL : リクエストの宛先

- ボディ : リクエストの本体
- `req.Body` でHTTPリクエストのボディを取得できる

```go
type Request struct {
    Method string
    URL *url.URL
    Body io.ReadCloser // リクエストボディを格納する
}
```

- `io.ReadCloser` 型（インターフェース）
    - `Read(p []byte) (n int, err error)`
        - ボディの中身を p に読み込む
    - `Close() error`
        - ボディの中身を読み終わったときにcloseする

- `errors.Is(err, io.EOF)` でボディの終端に達したか確認できる

### 2-6 jsonをGo構造体に変換する
- `func Unmarshal(data []byte, v any) error`
    - jsonデータをGo構造体に変換する関数
    - `v` は `&` をつける必要あり？


- `curl "URL" -X POST -d 'json'`
    - `-d` オプションでリクエストボディにデータを含めることができる

### 2-7 デコーダ・エンコーダ
- メモリとストリーム
    - メモリ: 必要な情報をまとめて扱う
        - `[]byte`
    - ストリーム: 必要な情報は少しずつ流れてくるものとして扱う
        - 標準入力とか

- `io.Reader` / `io.Writer` インターフェース
    - `io.Reader`
        - `Read(p []byte) (n int, err error)`: ストリームからデータを読み込んでバイトスライス（＝メモリ）に読み込む
    - `io.Writer`
        - `Write(p []byte) (n int, err error)`: バイトスライス（＝メモリ）からストリームにデータを書き出す

- `json.Decoder` / `json.Encoder`
    - `json.Decoder`
        - ストリームから得られるデータをGo構造体に変換する
        - `func NewDecoder(r io.Reader) *Decoder`
            - `r io.Reader` で指定されるストリームから流れてくるデータをjsonデコードするためのデコーダを作成する
            - `os.Stdin` や `req.Body` など
    - `json.Encoder`
        - Go構造体をストリームにjson形式で書き出す
        - `func NewEncoder(w io.Writer) *Encoder`
            - `w io.Writer` で指定されるストリームにjsonエンコードしたデータを書き出すためのエンコーダを作成する
            - `os.Stdout` や `w http.ResponseWriter` など

#### リファクタリング
- デコード処理
    - リファクタリング前
        - `req.Body.Read(buffer)`
        - バイトスライスにメモリを書き出す→バイトスライスからGo構造体に変換
    - リファクタリング後
        - `json.NewDecoder(req.Body).Decode(&reqArticle)`
        - ストリームから直接Go構造体に変換

```go
// リファクタリング前
buffer := make([]byte, req.ContentLength)
_, err := req.Body.Read(buffer)
if err != nil && !errors.Is(err, io.EOF) {
    http.Error(w, "Failed to read request body", http.StatusInternalServerError)
    return
}
var reqArticle models.Article
err = json.Unmarshal(buffer, &reqArticle)
if err != nil {
    http.Error(w, "Failed to unmarshal JSON", http.StatusBadRequest)
    return
}

// リファクタリング後
var reqArticle models.Article
err := json.NewDecoder(req.Body).Decode(&reqArticle)
if err != nil {
    http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
    return
}
```

- エンコード処理
    - リファクタリング前
        - `json.Marshal(article)`
        - Go構造体からバイトスライスに変換→バイトスライスをストリームに書き出す
    - リファクタリング後
        - `json.NewEncoder(w).Encode(article)`
        - Go構造体をストリームに直接書き出す

```go
// リファクタリング前
jsonData, err := json.Marshal(article)
if err != nil {
    http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
    return
}
w.Write(jsonData)
// リファクタリング後
err := json.NewEncoder(w).Encode(article)
if err != nil {
    http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
    return
}
```

### 2章まとめ
- Go構造体を使ってHTTPリクエストやレスポンスのモデルを定義する方法を学んだ
- jsonパッケージを使ってGo構造体とjsonデータを相互変換する方法を学んだ
- メモリとストリームの違いを勉強し、ストリームから直接Go構造体に変換したり、Go構造体をストリームに直接書き出したりする方法を学んだ

## 3
### 3-1 

```bash
docker-compose up
```

```bash
today@MacBook-Pro-10 db % mysql -h 127.0.0.1 -u docker samp
ledb -p < createTable.sql 
```
```bash
cat createTable.sql | docker exec -i db-for-go mysql -u<docker_user> -p<docker_user> sampledb && echo 'OK' || echo 'FAILED'
```
localhost上で動いているMySQLサーバー内にあるsampleデータベースに対して、dockerユーザーでcreateTable.sqlのクエリを実行する

```bash
docker exec -it db-for-go mysql -u<docker_user> -p<docker_pass> sampledb
```

```bash
mysql> show columns from articles;
+------------+------------------+------+-----+-------------------+----------------+
| Field      | Type             | Null | Key | Default           | Extra          |
+------------+------------------+------+-----+-------------------+----------------+
| article_id | int(10) unsigned | NO   | PRI | NULL              | auto_increment |
| title      | varchar(100)     | NO   |     | NULL              |                |
| contents   | text             | NO   |     | NULL              |                |
| username   | varchar(100)     | NO   |     | NULL              |                |
| nice       | int(11)          | NO   |     | 0                 |                |
| created_at | datetime         | NO   |     | CURRENT_TIMESTAMP |                |
+------------+------------------+------+-----+-------------------+----------------+
6 rows in set (0.00 sec)

mysql> show columns from comments;
+------------+------------------+------+-----+---------+----------------+
| Field      | Type             | Null | Key | Default | Extra          |
+------------+------------------+------+-----+---------+----------------+
| comment_id | int(10) unsigned | NO   | PRI | NULL    | auto_increment |
| article_id | int(10) unsigned | NO   | MUL | NULL    |                |
| message    | text             | NO   |     | NULL    |                |
| created_at | datetime         | YES  |     | NULL    |                |
+------------+------------------+------+-----+---------+----------------+
4 rows in set (0.03 sec)
```