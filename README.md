# 智铃校园

智铃校园核心系统包括 `Beego 后端 API 应用` + `TCP 服务端` + `Vue 单页应用` + `STM32f1 单片机程序`，另外提供了 `.bat批处理工具` 和 `智铃主控机模拟器`。

智铃校园后台管理系统的 B/S 客户端在本仓库中，C/S 客户端使用 `Electron` 开发，详情请见 [Bellex-Electron](https://github.com/Nomango/bellex-electron)

## 编译环境

建议使用的编译环境为:

- `Windows 10`
- `Golang 1.11`
- `Beego v1.11.1`
- `Bee v1.10.0`
- `Node.js 10.15.0`

### 模块和功能

- RESTful API 服务端（处理客户端的请求）
- TCP 服务端（处理智铃主控机的请求）
- RPC 服务端（实现 API 后端和 TCP 后端之间的通信）
- NTP 授时客户端

## 编译和运行

> 在本地环境下设置环境变量 `BELLEX_MODE=develope`

在项目根目录下打开 PowerShell ，并按下面的顺序编译前端、后端，然后启动。

#### 前端编译

进入 app 目录，执行批处理命令

```bat
cd app
npm install
npm run build
```

npm run build 的执行时间约 3 分钟左右，执行完毕后会将编译好的前端文件放置到 server/static 目录下

#### 后端编译

首先配置好数据库：

1. 在本地 MySQL 数据库中创建名为 bellex 的数据库
2. 在 conf/dev.yaml 文件中设置好数据库的连接方式
3. 后端启动时会自动加载该 yaml 文件中的配置项

在 PowerShell 运行批处理命令：

```bat
./make run server
```

启动 Beego 后端服务，启动后会自动创建数据表和管理员账号

打开新的 PowerShell 运行批处理命令：

```bat
./make run tcp
```

启动 TCP 服务端，启动后会等待 智铃主控机的 TCP 请求

### 测试 TCP 服务

在 PowerShell 运行批处理命令：

```bat
./make test tcp
```

将启动一个智铃主控机模拟器，模拟器可以模拟单片机向服务器请求的过程。

![](https://github.com/Nomango/bellex/blob/master/preview/test_tcp.png?raw=true)

### 测试 NTP 服务

在 PowerShell 运行批处理命令：

```bat
./make test ntp
```

向阿里云ECS、中科院、国家授时中心三个 NTP 服务发送请求，并自动解析 NTP 报文，将结果打印在控制台。

![](https://github.com/Nomango/bellex/blob/master/preview/test_ntp.png?raw=true)

### 更多

在 PowerShell 运行批处理命令：

```bat
./make help
```

查看更多使用介绍。
