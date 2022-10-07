FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
COPY httpd_vsftpd.sh /usr/local/bin
RUN yum install -y httpd vsftpd \
      && chmod 755 /usr/local/bin/httpd_vsftpd.sh \
      && useradd -m ftpuser01 \
      && echo "password1234" | passwd --stdin ftpuser01 \
      && ln -s /home/ftpuser01 /var/www/html/ftpuser01
      #  ln - make links between files
ENTRYPOINT [ "usr/local/bin/httpd_vsftpd.sh" ]