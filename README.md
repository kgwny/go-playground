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

## インデント修正
インデントでスペースが使われている場合にタブへ変換するコマンド例

```sh
gofmt -w your_file.go
```

特定のパス配下に存在するすべてのgoファイルを対象に含める場合のコマンド例
```sh
find /path/to/dir -name '*.go' -exec gofmt -w {} +
```
