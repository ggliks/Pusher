# Pusher 酱
Pusher 酱目前是一个基于 go-cqhttp 的安全知识 QQ 推送机器人。她的全部功能就是搜集棱角论坛、Seebug Paper、安全客的每日更新，然后推送到指定的 QQ 群里，同时还能搜集当天更新的CVE，对数据进行整理后以 Markdown 文件形式保存到 QQ 群文件内。

### Comming soon

- 增加钉钉群推送
- 增加企业微信群推送

## 截图

![](https://cdn.bingbingzi.cn/blog/20220114173218.png)

## 生成的文档

![](https://cdn.bingbingzi.cn/blog/20220114173307.png)
## 使用
在 Pusher 运行前，需要在本地开启一个 cqhttp 服务 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) 

### 这里列出一个简单的配置文件
```yaml
# go-cqhttp 默认配置文件
account:
  uin: 10086 # QQ账号
  password: '' # QQ密码
  encrypt: true
  status: 0
  relogin:
    delay: 3
    interval: 3
    max-times: 0
  use-sso-address: true

heartbeat:
  interval: 5

message:
  post-format: string
  ignore-invalid-cqcode: false
  force-fragment: false
  fix-url: false
  proxy-rewrite: ''
  report-self-message: false
  remove-reply-at: false
  extra-reply-data: false
  skip-mime-scan: false

output:
  log-level: warn
  log-aging: 15
  log-force-new: true
  log-colorful: true
  debug: false
  
default-middlewares: &default
  access-token: ''
  filter: ''
  rate-limit:
    enabled: false
    frequency: 1
    bucket: 1

database:
  leveldb:
    enable: true
  cache:
    image: data/image.db
    video: data/video.db
    
servers:
  - http:
      host: 127.0.0.1 # http 服务
      port: 5700 # http 服务端口
      timeout: 5
      middlewares:
        <<: *default
      post:
```

### 接着需要对 Pusher 进行配置

首先进入 Pusher 目录，创建一个文件`config.toml`，完成后赋予 **pusher** 一个可执行权限：

```shell
chmod +x pusher
```

接着执行 pusher 
```shell
./pusher
```
这就会将默认配置写入`config.toml`文件内。

## 配置文件结构
```toml
[setting]
  groupids = ["别名1,QQ群号1", "别名2,QQ群号2"]
  server = "http://127.0.0.1:5700"
  timeout = 5
```

需要手动修改`setting.groupids`节点中的信息，保存后再次运行 pusher，就会在每天九点半对指定 QQ 群进行推送。

## 测试

Pusher 提供了测试命令，可以快速测试机器人的可用性:

```shell
./pusher -t -g 别名,QQ群号
```

## 感谢以下开源项目
[https://github.com/Mrs4s/go-cqhttp](https://github.com/Mrs4s/go-cqhttp)

[https://github.com/Le0nsec/SecCrawler](https://github.com/Le0nsec/SecCrawler)

[https://github.com/EdgeSecurityTeam/EBot](https://github.com/EdgeSecurityTeam/EBot)

[https://github.com/binganao/dailyCVE](https://github.com/binganao/dailyCVE)
