# mkdir /root/apache-arg
# cd /root/apache-arg
# vi Dockerfile

FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
RUN yum install -y httpd iproute
RUN echo "Hello foreground Apache httpd." > /var/www/html/test.html
ENTRYPOINT ["/usr/sbin/httpd","-DFOREGROUND"]


# [--build-arg]で変数を指定する
# docker image build \
# -f ./Dockerfile \
# -t c75:apache01 \
# --build-arg http_proxy=http://proxy.your.site.com:8080 \
# --build-arg https_proxy=http://proxy.your.site.com:8080 \
