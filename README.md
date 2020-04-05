## UsedTimer
UsedTimer(中古の時計 -> よく壊れる)ということで命名
このアプリケーションで用いる時計は以下の仕様を満たす必要がある。

### 仕様について
1. 指定した周期毎に" "というメッセージをchannel経由で送信する
2. 指定した周期毎に時間切れとなり、errorを返して実行loopをbreakする
3. 指定した復活上限値までスレッドを再起動する(無限ループにすると都合が悪いため、このような仕様を採用)

## Usage
### golangがinstall済みの場合
`$GOPATH`配下にこのレポジトリを配置
> $ go run used_timer/main/main.go

### golangが未インストールでDockerがinstallされている場合
docker imageをbuild
> cd used_timer
> $ docker build -t golang/used_timer .

buildしたcontainerを起動
> $ docker run golang/used_timer