# mkdir /root/apache
# cd /root/apache
# vi Dockerfile.apache

# 使用するDockerイメージ名を指定
FROM centos:7.5.1804

# プロキシサーバーの設定
ENV http_proxy http://proxy.your.site.com:8080

# プロキシサーバーの設定
ENV https_proxy http://proxy.your.site.com:8080

# httpとiproute RPMパッケージをDockerイメージにインストール
RUN yum install -y httpd iproute && yum clean all

# テスト用のHTMLファイルを配置
RUN echo "Hello Apache." > /var/www/html/index.html

# コンテナ起動時にhttpdサービスが起動すうように設定
RUN systemctl enable httpd

# yum clean コマンドでキャッシュ削除・アップデートのチェックと実行
# https://easyramble.com/yum-clean-and-update.html

# yumコマンド
# https://atmarkit.itmedia.co.jp/ait/articles/1608/29/news019.html
