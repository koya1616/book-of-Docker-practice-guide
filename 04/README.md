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
