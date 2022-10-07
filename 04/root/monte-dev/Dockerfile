FROM centos:7.5.1804
ENV hhtp_proxy http://proxy.your.site.com:8080
ENV hhtps_proxy http://proxy.your.site.com:8080
# -y, --assumeyes Automatically answer yes for all questions.
RUN yum install -y gcc-gfortran libgfortran-static
COPY pi.f /tmp
WORKDIR /tmp
RUN gfortran -c -o pi.static.o pi.f
RUN gfortran \
      pi/static.o \
      /usr/lib/gcc/x86_64-redhat-linux/4.8.2/libgfortran.a \
      /usr/lib/gcc/x86_64-redhat-linux/4.8.2/libquadmath.a \
      -static-libgfortran -static-libgcc \
      -o pi.static
RUN mv /tmp/pi.static /usr/local/bin
CMD [ "/usr/local/bin/pi.static" ]