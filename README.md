# go-playground

## GO言語のインストール

```sh
brew install go
```

## バージョン確認
```sh
go version
```
go version go1.24.3 darwin/arm64

## 所在確認
```sh
which go
```
/opt/homebrew/bin/go

## 環境変数 GOPATH の設定

```sh
vim ~/.zshrc
```

以下の記述を追加
```sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```sh
source ~/.zshrc
```

```sh
go env GOPATH
```
