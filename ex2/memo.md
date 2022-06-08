# 2章　Go言語の基本文法
## 型
### 整数型
・int
    ->とりあえずこれを使っておけ。

その他の整数型
・int8
・int16
・int32
・int64

### 符号なし整数型
・uint
    -> マイナス値を持たない整数型。

その他の符号なし整数型
・uint8
・uint16
・uint32
・uint64

### その他の整数型
・uintptr
    -> ポインタを扱うための型。

・byte
    -> uint8の別名。

・rune
    -> int32の別名。javaのcharcter型。

### 不動小数点型
・float32
    -> javaのfloat型。

・float64
    -> javaのdouble型。

### 複素数
・complex64
・complex128

### 文字列
・string

### 真偽値
・bool

## リテラル
### 整数リテラル
・8進数　：先頭に「0」をつける
　　0123
・16真数：先頭に「0x」をつける
　　0x123


### 実数リテラル
・小数点をつける or 「e」をつける
 1.2
 1.0e3

### テキストリテラル
・rune  ：文字をシングルクォートで囲む
・string：文字をダブルクォートで囲む

### 真理値リテラル
「true」または「false」

## 変数の宣言
### 変数宣言
変数を1個定義する場合　　　　　　　　：「var 変数名 型」
    var val string
変数を複数定義する場合（同一型のみ）：「var 変数名1, 変数名2, ... 型」
    var val1, val2 string

### 代入
１変数の代入　：「変数名 = 値」
    val = "a"
複数変数の代入：「変数名1, 変数名2 = 値1, 値2」
    val1, val2 = "a", 1

### 暗黙敵型変換
型宣言を省略した書き方
　「変数名 := 値」
　  val := "a"

### キャスト
「型名(値)」
    var val int32 = 100
    var val2 int64 = int64(val)

### 文字列 <-> 数値
「strconv」パッケージ内の関数を使用する。
・文字列 -> 数値：「変数名1, 変数名2(エラー情報） = strconv.Atoi(値)」
    val, err = strconv.Atoi("1")
・数値 -> 文字列：「変数名1, 変数名2(エラー情報） = strconv.Itoa(値)」
    val, err = strconv.Itoa(1)

