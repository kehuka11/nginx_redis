# nginx_redis
nginxとRedisを使用した動的プロキシ


## 使い方

・sample_appのビルド

~~~shell
cd sample_app1
docker build -t sample_app1:latest .

cd sample_app2
docker build -t sample_app2:latest .
~~~

・Dockerコンテナ起動

~~~shell
docker-compose up -d
~~~
