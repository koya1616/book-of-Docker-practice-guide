#!/bin/bash

/usr/bin/pkill vsftpd

# httpdデーモンを起動
/usr/sbin/httpd -k restart

# vsftpdデーモンを起動
/usr/sbin/vsftpd /etc/vsftpd/vsftpd.conf

# 永久に終わらないtailコマンド
tail -f /dev/null