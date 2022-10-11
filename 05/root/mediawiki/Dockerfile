# Mediawikiを含むDockerイメージの作成

FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
RUN yum remove -y php* \
      && yum install -y \
      http://rpms.famillecollet.com/enterprise/remi-release-7.rpm \
      && yum install -y --enablerepo=remi,remi-php72 \
      php php-mysqlnd php-mbstring php-xml http openssh-clients \
      && curl -O \
      http://release.wikimedia.org/mediawiki/1.31/mediawiki-1.31.0.tar.gz \
      && tar xzvf /mediawiki-1.13.0.tar.gz -C /var/www/html/
WORKDIR /var/www/html
RUN ln -sf ./mediawiki-1.31.0 wiki \
      && systemctl enable httpd