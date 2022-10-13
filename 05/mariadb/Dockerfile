# MariaDBサーバーを含むDockerイメージの作成

FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
COPY MariaDB.repo /etc/yum.repos.d/MariaDB.repo
RUN yum install -y MariaDB-server MariaDB-client
COPY server.cnf /etc/my.cnf.d/
COPY mariadb.sh /mariadb.sh
RUN chmod +x /mariadb.sh
VOLUME ["/var/lib/mysql"]
EXPOSE 3306
ENTRYPOINT ["/mariadb.sh"]