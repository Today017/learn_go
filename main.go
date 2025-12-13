package main

import (
    "io"
    "log"
    "net/http"
)

func main() {
    //ハンドラ
    //HTTPリクエストを受け取って、それに対するレスポンス内容をコネクションに書き込む関数
    helloHandler := func(w http.ResponseWriter, req* http.Request) {
        //io.WriteString(w, s)
        //io.Writer型wに文字列sを書き込む
        io.WriteString(w, "Hello, world!\n")

        //io.Writer型=インターフェース
        /*
            type Write interface {
                Write(p []byte) (n int, err error)
            }
        */
        //-> メソッドとしてWrite(int n, error err)->[]byte pを持つならOK
        //C++のコンセプトみたいなやつ？
    }

    //http.HandleFuncにハンドラ関数を渡すことで、サーバーでこのハンドラが使われるようになる
    http.HandleFunc("/", helloHandler)

    //ターミナルへのログ表示
    log.Println("server start at port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
    //http.ListenAndServe(addr string, handler Handler) error
    //addr=アドレス
    //error=サーバー起動時に起こったエラー
    //errorをlog.Fatelに渡している
    //log.Fatal=エラーをログ出力して異常終了

    //handler=nilは？
}
