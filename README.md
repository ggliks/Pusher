# Pusher 酱
Pusher 酱是一个基于 go-cqhttp 的安全知识 QQ 推送机器人

## 截图

![](https://cdn.bingbingzi.cn/blog/20220114173218.png)

![](https://cdn.bingbingzi.cn/blog/20220114173307.png)
## 使用
在 Pusher 运行前，需要在本地开启一个 cqhttp 服务 [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) ，这里列出一个简单的配置文件:

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
