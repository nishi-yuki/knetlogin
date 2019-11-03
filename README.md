# knetlogin
k**netの自動ログインツール

## インストール
goの開発環境がある場合
```shell
go get github.com/Ni5h1/knetlogin
```

## 使い方

k**netに接続した状態で、
```shell
$ knetlogin -i {your id} -p {your password}
```
あるいは
`knetlogin.yaml`(後述)にidとパスワードを記述して、
```shell
$ knetlogin
```
を実行すると、k**netにログインできます。
なお、外部のインターネットに接続する際はプロクシの設定が別途必要です。

### knetlogin.yaml
`knetlogin.yaml`にidとパスワードを以下の形式で記述します。

```knetlogin.yaml
id: your id
pass: your password
```
`knetlogin.yaml`は下の箇条書されたいずれかの場所に保存してください。
複数存在した場合、より上に書かれたものが優先されます。
1. `~/.config/knetlogin.yaml`
1. 実行ファイルがあるディレクトリ
