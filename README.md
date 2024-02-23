# fib_api
###  1. 概要
###  2. ファイル構成 
  * main.go - TCPコネクションからserverのｈａｎｄｌｅｒを呼び出す
  * middleware.go - corsを設定し、get・option以外の呼び出しを拒否している
  * routes.go -　routeを設定
  * utils.go - json → object, object → json
  * handlers.go - 引数をもとにフィボナッチ数を出力するhandlerを格納
  * handlers_test.go - handlers.goファイルの関数をテストする(正常系、異常系)
  * go.mod - moduleの定義、サードパッケージの管理
### 3. 工夫した点
  1. フィボナッチ数列をメモ化を用いて再現することで高速化している
  2. 
