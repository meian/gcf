# gcf

gcf(Go Colletion Framework) は Generics を用いて様々なコレクション操作を提供するライブラリです。  
コレクションに対する操作を共通のインターフェイスを用いて行うことにより、操作の合成を容易に行うことができるようになります。

[![Go Reference](https://pkg.go.dev/badge/github.com/meian/gcf.svg)](https://pkg.go.dev/github.com/meian/gcf)
[![codecov](https://codecov.io/gh/meian/gcf/branch/main/graph/badge.svg?token=PDHAVSGE0E)](https://codecov.io/gh/meian/gcf)
[![Go Report Card](https://goreportcard.com/badge/github.com/meian/gcf)](https://goreportcard.com/report/github.com/meian/gcf)


## モチベーション

Goでもforやifなどの基本構文での処理でなく、複数の処理を容易に合成して利用できる仕組みが欲しいと考えました。

これまでは同じインターフェイスで複数の型に処理を提供することが困難でしたが、Go 1.18においてGenericsがサポートされたことでこの実装が容易になったため、実際にライブラリとして構築したものがgcfです。

## Example

`スライスの要素のうち奇数の数値のみを抽出し、それらの数値を3倍して返す` という処理を例として取り上げます。

```golang
func Odd3(s []int) []int {
    var r []int
    for _, v := range s {
        if v%2 == 0 {
            continue
        }
        r := append(v*3)
    }
    return r
}
```

gcfを用いた場合は以下のように実装します。

```golang
// var s []int

itb := gcf.FromSlice(s)
itb = gcf.Filter(itb, func(v int) bool {
    return v%2 > 0
})
itb = gcf.Map(itb, func(v int) int {
    return v * 3
})

// 処理結果をスライスで取得する
r := gcf.ToSlice(itb)
```

この例は利用方法を簡潔に示すためのものです。  
インラインの処理をgcfに置き換えると著しく処理速度が低下するため、インラインで容易に実装・管理できる処理をgcfを用いて書き換える事はお勧めしません。

## 環境

- Go 1.18 or 1.19

gcf では Generics を利用しているため、1.18以上のバージョンが必要になります。  
ローカル環境に新しいバージョンのGoをインストールしたくない場合に利用できるvscodeでの利用に合わせたコンテナ利用環境も用意しています。  
([.devcontainer](https://github.com/meian/gcf/tree/main/.devcontainer) 以下を参照)

## インストール方法

Goのモジュールで管理されてる配下のディレクトリ上で `go get` を用いることでインストールします。

```bash
go get -d github.com/meian/gcf
```

## デザイン

### Iterator による実装

gcfは `Iterator` パターンによって処理を連携するよう構成されています。  
いくつかの処理では内部で追加のメモリアロケーションを行うこともありますが、大半の処理では処理途中で不必要なメモリのアロケーションが発生しないようになっています。

### Iterable + Iterator

gcf の各関数は `Iterable[T]` のインターフェイスを持ち、これは `Iterator()` によって `Iterator[T]` を生成する機能のみを持ちます。
`Iterator[T]` は `MoveNext()` によってコレクションから取得できる要素を次の要素に移動し、`Current()` によって現在の位置の要素を取得します。  
`Iterable[T]` によって操作を合成し、状態はそこから生成される `Iterator[T]` にのみ保持されることで、生成した操作を再利用しやすくなっています。

### MoveNext + Current

Iterator パターンの実装では `HasNext()` で次の要素があるかを確認し、 `Next()` で次の要素に移動して要素を返す処理の組み合わせの実装が見られることがあります。  
gcf では `MoveNext()` で次の要素に移動してその成否を返し、`Current()` で現在の要素を返すといういう実装を採用しました。  
これは、状態を変更しないで次の要素があるかどうかを確認する機能を提供するメリットよりも、状態を変更しないで値を何度も取得できるメリットを採用したかったためです。

### トップレベル関数

他の言語にあるコレクション操作のライブラリではメソッドチェーンによって扱えるように処理をメソッドで定義することが多いですが、gcfで提供されるコレクション処理はトップレベルの関数で実装されます。  
これはGoにおけるGenericsではメソッドレベルでのタイプパラメータを定義することができないため、メソッドでの機能提供を行えない機能が出てきてしまい、一部機能のみをメソッド提供すると処理に一貫性を保てないためです。  
Goの今後のバージョンアップによりメソッドレベルでもタイプパラメータの定義ができるようになれば、メソッドチェーンの機能提供も検討します。

### スレッドセーフ

Iterableを生成するexportedな関数はスレッドセーフで、複数のgo routineから同時にアクセスすることが出来ます。  
それぞれのIterableにおける `Iterator()` の処理もスレッドセーフで、複数のgo routineから同時に呼び出しが可能です。  
それぞれのIteratorにおける `MoveNext()` や `Current()` の処理はスレッドセーフ性は保証されないため、go routine間で共用する場合は必要に応じて mutexなどで処理の分離を行ってください。  
`ToSlice()` などのIterableから内部要素を取り出す処理はスレッドセーフです。  

現在は未実装ですが、channelに関する実装が追加された際にスレッドセーフの特性がどのように変更されるかについては未定です。

## パフォーマンス

gcfのパフォーマンスは以下の特性を持っています。

- コレクションの要素数や合成する処理数に比例した処理時間がかかる
- インラインの処理と比較すると圧倒的に遅い(70倍程度)
- アロケーションなしの関数呼び出しと比較すると4倍程度遅い
- チャンネル処理と比較するとと圧倒的に早い(60倍程度)

ライブラリの特性上繰り返し処理されるものであるため、シビアな処理速度を求められる処理に対しての利用はお勧めできません。

詳細は [ベンチマークのREADME](bench/README.ja.md) にまとめたのでそちらを参考にしてください。

## 今後検討中の機能

- channel系の機能
  - channelからIterableを作成
  - Iterableの結果をchannelで取得
  - 設計イメージが思い浮かばないので休止中
- `Zip`
  - 複数のIterableの要素を合成して一つのIterableにまとめる
  - それぞれのIterableの要素数が違う場合はどう処理するかは他の実装を参考にする予定

----

## READMEの表示言語を構成する

- [English](README.md)