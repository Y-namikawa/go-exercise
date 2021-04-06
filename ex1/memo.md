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

# 標準入力
```
package main

import (
  "bufio"
  "os"
)

func main() {
  // 入力用の変数を定義
  scanner := bufio.NewScanner(os.Stdin)
  // コマンドラインから入力を受け付ける
  scanner.Scan()
  // 入力された文字列を取得する
  scanner.Text()
}
```

＃GOPATH
- Go言語が参照する場所
- Goをインストールしたフォルダーないの"src"を示す
- "src"内に作成したプログラムは他のプログラムでimportできる
