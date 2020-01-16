# intab-ops

## git clone
``` shell
git clone --recursive git@code.aliyun.com:intab/intab-ops.git
```

## dockerfiles

``` shell
Dockerfile.intab-apiserver              #intab-apiserver 生产环境镜像

Dockerfile.intab-apiserver.dev          #intab-apiserver 开发环境镜像

Dockerfile.intab-webserver              #intab-webserver 生产环境镜像

Dockerfile.intab-webserver.dev          #intab-webserver 开发环境镜像

Dockerfile.intab.all                    #开发环境基础镜像，将依赖的第三方包集成进来

Dockerfile.intab.all.local              #开发环境基础镜像，将依赖的第三方包集成进来，网络不好时，从之前的镜像中提取

```

## deploy

``` shell
deploy.all.sh                           #生成开发环境基础镜像
deploy.all.dev.sh                       #生成开发环境基础镜像，网络不好时使用
deploy.dev.sh                           #生成各种开发环境镜像
deploy.sh                               #生成各种生产环境镜像
```

## docker-compose.yml

``` shell
#启动开发&测试环境
docker-compose -f docker-compose.aben.yml up -d

#更新某开发镜像
docker rm -f intab-apiserver-dev
docker-compose -f docker-compose.aben.yml up -d

#服务运行异常时，请尝试彻底更新环境
docker-compose -f docker-compose.aben.yml down
docker-compose -f docker-compose.aben.yml up -d
```

## 容器管理

``` shell
#登录
docker login --username=docker-admin@intab registry.cn-beijing.aliyuncs.com

docker tag [ImageId] registry.cn-beijing.aliyuncs.com/intab/intab-all:[镜像版本号]
docker push registry.cn-beijing.aliyuncs.com/intab/intab-all:[镜像版本号]
```
