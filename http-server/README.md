# HTTP Server 

基本的なHTTPサーバにはいくつかの仕事がある

- ウェブサイトをブラウズしているユーザから送信されるリクエストを処理し、アカウントにログインするかイメージを投稿する
- ダイナミックなユーザ体験を提供するために、JSやCSS、画像などをブラウザに送信する
- 指定したポートをリッスンしコネクションを受け入れる

## Process dynamic requests

`net/http`パッケージは、requestを受付けて動的に処理するなど必要とされる機能をすべて含んでいる。`http.HandleFunc`関数で新しいハンドラを登録できる。第一引数にはマッチするパスを、第二引数には実行する関数を取ります。もしウェブサイト（http://example.com/）にアクセスしたときに、ユーザに挨拶する場合には、下記のようになります。

```go
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Welcome to my website!")
})
```

`http.Request`はrequestとそのパラメータについての情報をすべて持っています。そのため、GETのパラメータは`r.URL.Query().Get("token")`で取得できるし、POSTのパラメータ（HTMLのform）は`r.FormValue("email")`で取得できる。

## Serving static assets

JSやCSS、画像などを配信する際には、ビルトインの`http.FileServer`を使い配信対象のURLとパスを示します。

```go
fs := http.FileServer(http.Dir("static/"))
```

ファイルサーバを配置したら、URLパスを指定するだけです。  
※ ファイルを正しく配信するために、URLパスの一部を取る必要があります。通常、その名前は、配信対象のファイルが配置されたディレクトリ名となります。

```go
http.Handle("/static/", http.StripPrefix("/static/", fs))
```

## Accept connections

HTTPサーバの公開

```go
http.ListenAndServe(":80", nil)
```
