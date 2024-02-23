# fib_api
###  1. 概要
  #### 引数からフィボナッチ数を出力するAPI作成
    ・url: https://fib-api-service.onrender.com
    ・使用例: https://fib-api-service.onrender.com/?n=3
    　クエリに n = 整数 を持たせることでnに基づいたフィボナッチ数を出力します。
     ex) n = 3 → result = 2,
         n = 18 → result = 2584,
         n = 33 → result = 3524578,
  
###  2. ファイル構成 
  * main.go - TCPコネクションからserverのhandlerを呼び出す
  * middleware.go - corsを設定し、get・option以外の呼び出しを拒否している
  * routes.go -　routeを設定
  * utils.go - json → object, object → json
  * handlers.go - 引数をもとにフィボナッチ数を出力するhandlerを格納
  * handlers_test.go - handlers.goファイルの関数をテストする(正常系、異常系)
  * go.mod - moduleの定義、サードパッケージの管理
### 3. 工夫した点
  1. フィボナッチ数列をメモ化を用いて再現することで高速化している
  2. json周りの処理をメソッド化することで、コード量を減らしファイルを見やすいようにしている
  3. routerはgoの標準パッケージを使うのではなく、go-chiというサードパッケージを使うことでmiddlewareの設定を簡単にできるようにしている。

### ※注意点
* 最初の通信に20秒ほどかかってしまいます。
* goの型の上限の問題で、n =　1 ~ 93までの整数しかクエリとして受け入れることができません。
  →　goはuint64で最大64ビット格納できるが、整数18,446,744,073,709,551,615までの数しか対応できない。そのため、94以上だとこの数値を超えてしまい誤った数値を出力してしまう。
