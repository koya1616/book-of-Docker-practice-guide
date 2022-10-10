## Dockerfile

#### WebサーバーのDockerfile作成からコンテナ起動までの行程
1. Dockerfile作成
2. Dockerイメージのビルド
  ```
  % docker image build -f ./Dockerfile -t centos:c75apache01 .
  ```
3. Dockerコンテナの起動
  ```
  % docker container run -itd --tmpfs /tmp --tmpfs /run \
    -v /sys/fs/cgroup:/sys/fs/cgroup:ro --stop-single SIGRTMIN+3 \
    --name web0001 -h web0001 -p 8081:80 centos:c75apache01 /sbin/init

  roをつけることでread onlyでボリュームを共有できる
  停止シグナルを受け付けるように設定する必要がある。 アプリケーションで定義した用途に使用できるシグナルSIGRTMIN+3を設定する
  ```
4. 稼働中のDockerコンテナにログイン
  ```
  % docker container exec -it web0001 /bin/bash
  ```

#### 【bash入門】bashシェルスクリプトの書き方
- https://tech-blog.rakus.co.jp/entry/20210525/shellscript

#### 死活監視を行うオプションを付与したコンテナの実行
```
% docker container run \
    --rm \
    -itd \

    <!-- --name c75-httpd-wget \ -->
    --name c75-httpd \

    -p 81:80 \

    <!-- --health-cmd='wget -q -0 - http://localhost/test.html || exit 1' \ -->
    --health-cmd='curl -f -s -o /dev/null http://localhost/test.html || exit 1' \

    --health-interval=1s \
    --health-timeout=1s \
    --health-retries=3 \

    <!-- c75:httpd-wget -->
    c75:httpd
```

#### Dockerfileの利用指針
- 不要なファイルやディレクトリを置かない
  - docker image buildを行う際に、ファイルやディレクトリも含めてDockerのコンテキストに追加してしまう。
- 「.dockerignore」ファイルを使用する
- 変更のない手順については、既存のキャッシュを使い、変更部分のみをビルドすることで、全体のビルド時間を短縮する。
  - yum update すべきパッケージが公開されてたとしても、変更がなければupdateされない
  - yumの場合は「--no-cache=true」を使う？
- アップデートするパッケージとバージョンは明示的に指定するべき
- cd コマンドを使わない
  - WORKDIR命令を使おう
- RUN命令の数を削減する
  - RUN命令を複数記述すると、生成されるDockerイメージのサイズが肥大化します。