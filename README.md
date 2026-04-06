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
# これは多分エラーになるので以下のコマンドで代用

docker exec -it db-for-go mysql -u docker -p
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


### 3-7

トランザクション：「複数個のクエリをまとめて一つの処理として扱う」

1. トランザクションを張る
2. クエリを実行する
3. 全部成功したら、コミットして結果を確定させる
4. どれかが失敗したら、ロールバックしてなかったことにする


## 4
### 4-2

テーブルドリブンテスト

```go
tests := []struct {
    testTitle string
    exptected models.Article
}{
    {
        testTitle: "subtest1",
        exptected: models.Article{
            ...
        },
    }, {
        testTitle: "subtest2",
        exptected: models.Article{
            ...
        },
    },
}

for _, test := range tests {
    t.Run(test.testTitle, func(t *testing.T) {
        got, err := repositories.SelectArticleDetail(db, test.exptected.ID)
        if err != nil {
            t.Fatal(err)
        }

        ...
    })
}
```

### 4-3

前処理と後処理を共通化する。
グローバル変数を定義して `setup()` `teardown()` に前処理を書く。

`TestMain(m *testing.M)` という関数を定義して一連の処理の流れを書く。

```go
func TestMain(m *testing.M) {
	setup()

	m.Run()

	teardown()
}
```

### 4-4

個別のテストで後処理を書きたい場合
- insertのテストをしたあとそのデータを消す、など

`testing.T::Cleanup()` を使って、テスト後に行いたい処理を書く。

```go
t.Cleanup(func() {
    ...
});
```

`defer func() {...}()` ではだめなのか？
- 以下のような書き方をするとサブテストが並列で回るようになるが、このときに `defer` だと後処理→サブテストの順番で走ってしまう

```go
t.Run(testcase.testTitle, func(t *testing.T) {
    t.Parallel() // これがあることで、サブテストが並列に走るようになる
    fmt.Println(testcase.testTitle)
})
```

### 4-5


#### 概要
Goのテスト実行時に、MySQL（Docker）のテストデータを初期化する処理（`setupDB.sql`, `cleanupDB.sql`）が失敗する問題が発生。
複数の原因が絡み合っていたため、エラー出力を可視化しながら段階的に原因を特定・解消した。

---

#### バグ1：ターミナルからSQLが実行できない（認証プラグインエラー）

##### 🔍 発見の経緯
ターミナルから直接 `mysql` コマンドでファイル流し込み（`<`）を試したところ、以下のエラーが直接出力された。
> `ERROR 2059 (HY000): Authentication plugin 'mysql_native_password' cannot be loaded`

##### 💣 原因
- **環境の不一致**: Mac側のMySQLクライアント（Homebrew）が **v9.5.0** と新しく、MySQL 9.0で廃止された古い認証方式（`mysql_native_password`）に非対応だった。
- サーバー側（DockerのMySQL）は古い認証方式を求めてきたため、クライアント側がパニックを起こしていた。

##### 💊 解決策
Macローカルの `mysql` コマンドを使うのをやめ、**`docker exec -i` を経由してコンテナ内部の（バージョンが一致している） `mysql` コマンドを直接叩く**アプローチに変更した。

---

#### バグ2：Goから実行すると `exit status 1` で落ちる（エラーの隠蔽）

##### 🔍 発見の経緯
バグ1の解決策をGoの `exec.Command` に組み込んで `go test` を実行したが、詳細な理由が表示されずただ **`exit status 1`** とだけ出力されてテストが落ちた。

##### 💣 原因
Goの `exec.Command` は、デフォルトでは**外部コマンド（今回ならMySQL）の標準エラー出力（`Stderr`）をターミナルに表示せず飲み込んでしまう仕様**になっている。
そのため、実際にはMySQL側で何らかのエラーが起きているのに、Go側からは「とにかく失敗した」ということしか分からない状態になっていた。

##### 💊 解決策
原因を特定するため、Goのコードを修正し、**MySQLの標準エラー出力をキャッチして画面に表示する**仕組み（`bytes.Buffer`）を追加した。（※これが最大の突破口！）

---

#### バグ3：隠れていたSQLの構文エラー（タイポ）

結局原因はこれだった

---

#### 最終的な実装（修正完了版）

これらの知見を踏まえ、今後の拡張も考慮して「指定したSQLファイルをDocker経由で実行し、エラーを逃さず表示する共通関数」を作成した。

```go
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// execSQLFile は指定されたSQLファイルをDockerコンテナ内のMySQLで実行する
func execSQLFile(filePath string) error {
	// 1. ファイル読み込み
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 2. Docker経由でMySQLコマンドを準備（Macのバージョン依存を回避）
	cmd := exec.Command(
		"docker", "exec", "-i", "db-for-go",
		"mysql", "-u", "docker", "--password=docker", "sampledb",
	)

	// 3. 標準入力とエラー出力を設定
	cmd.Stdin = file            // ファイルの内容を流し込む
	var stderr bytes.Buffer
	cmd.Stderr = &stderr        // MySQLの詳細なエラーメッセージを捕まえる

	// 4. 実行とエラーハンドリング
	if err := cmd.Run(); err != nil {
		// 捕まえたMySQLのエラー詳細をGoのエラーとして返す
		return fmt.Errorf("SQL実行エラー: %v\n【詳細】: %s", err, stderr.String())
	}
	return nil
}
```

#### 💡 学んだこと・教訓
* 外部コマンド呼び出しで `exit status 1` のような抽象的なエラーが出た時は、**真っ先に `cmd.Stderr` をキャッチしてログに出す**。これがデバッグの基本にして最大の近道。
* 認証プラグインエラーなどの環境依存バグと、タイポなどの単純なバグが同時に発生すると混乱しやすい。ログを出力して**「今どこで（Mac? Docker? Go? SQL?）エラーが起きているか」を切り分ける**ことが重要。

## 5

レポジトリ層から得たデータをハンドラ層が必要とする形に加工して2つの層の橋渡しをする役をサービス層という。



>例外処理で返す値について注意！
>構造体を返すときは空の構造体で良いが、構造体のスライスを返すときは `nil` を返せるので、長さ 0 のスライスと区別するために `nil` を返すようにする

>スライスで `nil` を返せる理由
>構造体や string, int などは、「値型」といって、変数を作った瞬間にメモリが確保される様になっている。
>よって、 `nil` になることは物理的に不可能。
>一方で、スライスやマップは「スライス型」といって、アドレスを指し示す矢印となっている。
>アドレスをがどこも指していない状態である `nil` を取ることができる。

## 6

## 6-1

アーキテクチャで良くない部分
- `sql.DB` をOpen/Closeする頻度が多い
- サービス層の中にデータベース接続処理が内包されている
    - サービス層のコードがデータベース接続処理に依存して動いている
        - ブログAPIの根幹部分たるサービス層がデータベースという不安定な外聞要因に依存しているのは良くない
        - データベースをMySQLからPostgresSQLに変更したらサービス層も大幅変更を強いられる
        - サービス層の役割は「ハンドラ層用のデータを整形する」こと
- データベースの情報が `services/helper.go` にあるというのがわかりにくい
    - なんのデータベースを使っているかという情報はインフラの観点でかなり重要な情報だが、それがディレクトリの奥深くの1ファイルにかかれているのは不親切

サービス構造体を導入する
