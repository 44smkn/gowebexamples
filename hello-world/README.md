# Hello World

[https://gowebexamples.com/hello-world/](https://gowebexamples.com/hello-world/)

## RequestHandlerの登録

Handlerに当たるのが、

```go
func (w http.ResponseWriter, r *http.Request)
```

です。  
`http.ResponseWriter`は、text/htmlのレスポンスを書き込む場所を指し、`http.Request`はrequestのすべての情報を含む（例えばURLやヘッダーフィールドなど）

## HTTP Connectionのリッスン

RequestHandler単独では、外部からHTTPコネクションを受け付けることはできません。HTTP ServerはRequestHandlerにコネクションを渡すためにポートをリッスンする必要性があります。ほとんどのケースで80番ポートをリッスンするので、今回もそれに習います。

```go
http.ListenAndServe(":80", nil)
```