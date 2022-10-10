# SHELLによるシェルの変更

FROM centos:7.5.1804
ARG http_proxy
ARG https_proxy
COPY dot.bash_profile /root/.bash_profile
COPY dot.tcshrc /root/.tcshrc
RUN yum install -y tcsh

# Override the default shell used for the shell form of commands.
SHELL [ "/bin/bash", "--login", "-c" ]

RUN echo $0     >> /root/shell.log && \
    echo $LANG  >> /root/shell.log && \
    echo $PATH  >> /root/shell.log && \
    echo $VAL01 >> /root/shell.log && \
    echo ""     >> /root/shell.log
SHELL [ "/bin/bash", "-c" ]
RUN echo $0     >> /root/shell.log && \
    echo $LANG  >> /root/shell.log && \
    echo $PATH  >> /root/shell.log && \
    echo $VAL01 >> /root/shell.log && \
    echo ""     >> /root/shell.log
ENTRYPOINT [ "tail", "-f", "/dev/null" ]