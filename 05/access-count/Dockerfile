# Webアクセス数をカウントするアプリケーションの作成

FROM python:3.4-alpine
ARG http_proxy
ARG https_proxy
ADD . /code
WORKDIR /code
RUN pip install flask \
      && pip install redis
CMD [ "python", "ac.py"]