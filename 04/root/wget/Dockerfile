# システム要件や運用上の理由から、curlコマンドが使えない場合もあります。
# そのような場合は、wgetコマンドを使った死活監視が可能です。

FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
RUN yum install -y httpd wget && \
      echo "Hello Docker." > /var/www/html/test.html
ENTRYPOINT [ "/usr/sbin/httpd", "-DEFOREGROUND" ]