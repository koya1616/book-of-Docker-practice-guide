## Docker Community Edition

#### Docker環境のためのホストOSのインストール
1. 物理サーバーへのCentOSのインストール
2. ディスクのパーティショニング
3. パッケージのインストール
4. OSのアップデート
5. Dockerのインストール
6. Dockerのパラメータ設定
7. Dockerの状態確認

#### docker image lsのオプション
```
本来の６４文字のイメージIDを表示
% docker image ls --no-trunc

imageIDのみ出力
% docker image ls --no-trunc -q
```

#### リソースの使用状況の確認
```
% docker container stats

-a: すべてのコンテナ
--no-stream: 一度だけ表示する
```
