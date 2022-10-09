# マルチステージ用のDokerfile

# the name of this build stage.
# set the base image to use for any subsequent instructions that follow and also give this build stage a name.
# 開発用のDocekrイメージ。
FROM centos:7.5.1804 AS build0001

ENV hhtp_proxy http://proxy.your.site.com:8080
ENV hhtps_proxy http://proxy.your.site.com:8080
RUN yum install -y gcc-gfortran libgfortran-static libquadmath-static
COPY pi.f /tmp
WORKDIR /tmp
RUN gfortran -c -o pi.static.o pi.f
RUN gfortran \
      pi/static.o \
      /usr/lib/gcc/x86_64-redhat-linux/4.8.2/libgfortran.a \
      /usr/lib/gcc/x86_64-redhat-linux/4.8.2/libquadmath.a \
      -static-libgfortran -static-libgcc \
      -o pi.static

# 本番環境用のDockerイメージ
FROM centos:7.5.1804 AS run0001

# 開発用のイメージからpi.staticをコピー
COPY --from=build0001 /tmp/static /usr/local/bin
CMD [ "/usr/local/bin/pi.static" ]