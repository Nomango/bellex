# bellex

智铃校园

### 模块和功能

- NTP 授时服务
- TCP 服务端，提供校时服务

### 环境

> `Windows 10` + `Golang 1.11`

### 测试 NTP 服务

在 PowerShell 运行批处理命令：

```bat
./make test_ntp
```

将自动向阿里云ECS、中科院、国家授时中心三个 NTP 服务发送请求，并自动解析 NTP 报文，将结果打印在控制台。

![](https://github.com/Nomango/bellex/blob/master/preview/test_ntp.png?raw=true)

### 测试 TCP 服务

在 PowerShell 运行批处理命令：

```bat
./make test_tcp
```

将启动一个 TCP 连接，每隔 1 秒发送一次请求，共发送 5 次。服务器会返回一个时间戳，以完成校时。

![](https://github.com/Nomango/bellex/blob/master/preview/test_tcp.png?raw=true)
