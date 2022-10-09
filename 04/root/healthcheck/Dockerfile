FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
COPY check.sh /
RUN install -y httpd && \
      echo "Hello Dokcer." > /var/wwww/html/test.html && \
      chmod +x / check.sh
HEALTHCHECK \
  --interval=1s \
  --timeout=1s \
  --retries=3 \
CMD /check.sh || exit 1
ENTRYPOINT [ "/usr/sbin/httpd","-DFOREGROUND" ]