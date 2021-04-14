# 重定向器
用于重定向80端口到https，适用于 Xray监听443端口，重定向器监听80转发，且有个静态网站
* 静态网站服务器



## 使用
命令行参数：
```
  -config string
        Path to config file (default "./config.yaml")
```

<details>
  <summary>点击此处展开示例配置文件</summary>
  
```yml
# listen: 监听地址
listen: 127.0.0.1:8080

# redirecthttps: 监听一个地址，发送到这个地址的 http 请求将被重定向到 https
redirecthttps: 0.0.0.0:80

# inboundbuffersize: 入站缓冲区大小，单位 KB, 默认值 4
# 相同吞吐量和连接数情况下，缓冲区越大，消耗的内存越大，消耗 CPU 时间越少。在网络吞吐量较低时，缓存过大可能增加延迟。
inboundbuffersize: 4

# outboundbuffersize: 出站缓冲区大小，单位 KB, 默认值 32
outboundbuffersize: 32

# http: listen 的 http 流量的处理方式
http:
  # handler: fileServer 将服务一个静态网站
  handler: fileServer
  # args: 静态网站的文件路径
  args: /var/www/html

#  # handler: proxyPass 将流量转发至另一个地址
#  handler: proxyPass
#  # args: 转发的目标地址
#  args: 127.0.0.1:40001

```
</details>
