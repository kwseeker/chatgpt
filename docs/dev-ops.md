# DevOps



## 改版工程运维



## 原版工程运维

### Portainer

[Install Portainer CE with Docker on Linux](https://docs.portainer.io/start/install-ce/server/docker/linux)

是Docker容器运维工具，本身也按容器的方式启动，TCP tunnel server 端口 8000，WebUI HTTPS 端口 9443, HTTP 端口 9000。

 Portainer 提供*Portainer Server* 和 *Portainer Agent*（比如Portainer容器部署在云服务器，本地安装Agent远程运维）。

```shell
docker run -d -p 8000:8000 -p 9443:9443 -p 9000:9000 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:latest
```

### Nginx

```shell
# docker pull nginx:stable-alpine3.17
# 本地开发环境不需要	--restart always
docker run \
	-d \
	-p80:80 \
	-p443:443 \
	--name nginx-single \
	-v /data/nginx/ssl/openai.kwseeker.top.pem:/etc/nginx/ssl/openai.kwseeker.top.pem \
	-v /data/nginx/ssl/openai.kwseeker.top.key:/etc/nginx/ssl/openai.kwseeker.top.key \
	-v /data/nginx/conf.d/openai.kwseeker.top.conf:/etc/nginx/conf.d/openai.kwseeker.top.conf \
	nginx:stable-alpine3.17
```

支持TLS：

```shell
# openssl 本地生成证书，生产环境是去服务商申请的,包含两个文件 ***.key ***.pem
# 参考：https://www.baeldung.com/openssl-self-signed-cert
openssl genrsa -des3 -out openai.kwseeker.top.key 2048
# 输入PEM密码：比如123456
openssl req -x509 -newkey rsa:4096 -sha256 -nodes -keyout openai.kwseeker.top.key -out openai.kwseeker.top.pem -days 3650
# 然后输入证书信息回车即可生成
# 最后拿到 openai.kwseeker.top.key、openai.kwseeker.top.pem
```

最后配置Nginx配置文件 openai.kwseeker.top.conf。

