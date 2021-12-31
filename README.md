# golang の導入

1. `sudo add-apt-repository ppa:longsleep/golang-backports`
2. `sudo apt update`
3. `sudo apt install golang`

# echo の導入手順

1. `go mod init <app名>`
2. `go get -u github.com/labstack/echo/v4`

# 構造体関係で参考になりそうなサイト

- [構造体の基礎](https://golang.hateblo.jp/entry/golang-how-to-use-struct)
- [構造体のメソッドによる処理の紐づけ](https://www.wakuwakubank.com/posts/779-go-struct-method-interface/)

### 注意点

- [package 外出し](https://qiita.com/zurazurataicho/items/4a95e0daf0d960cfc2f7)
- [コンストラクタへの代入](https://netmark.jp/2017/08/post-1400.html)
