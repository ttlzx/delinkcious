FROM ttlzx/delinkcious-python-flask-grpc:0.1
MAINTAINER Gigi Sayfan "the.ttlzx@gmail.com"
COPY . /api_gateway_service
WORKDIR /api_gateway_service
EXPOSE 5000
ENTRYPOINT FLASK_DEBUG=1 python run.py
