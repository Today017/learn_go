# learn_go

[APIを作りながら進むGo中級者への道](https://techbookfest.org/product/jXDAEU1dR53kbZkgtDm9zx?productVariantID=dvjtgpjw8VDTXNqKaanTVi)

2025/12/13勉強開始

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

