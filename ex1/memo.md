# 1章.Go言語をスタートする
## Go言語の基本構文
```
package パッケージ名

import (
    読み込むパッケージ名
)

func main() {
  処理
}
```
- 変数名は半角英数字で書く
- 大文字・小文字は区別される
- インデントは不要
- 文末に;(セミコロン)が不要
- 演算子の前後にスペースを入れないのが基本。スペースは計算の優先度の定義にしようするため

## Go言語でHello World!
```
package main

import (
    "fmt"
)

func main() {
  fmt.Println("Hello World!")
}
```
