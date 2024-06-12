# GOGOGO

弥补CloudweGo课程的遗憾，学习GO语言。

## Get Start

开发环境：`阿里云Ubuntu20.04`

### 安装配置GO语言环境

- 访问[go.dev/dl](https://go.dev/dl/)并下载压缩包。
- 执行安装程序：`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz`
- 修改`/etc/profile`配置环境变量和代理：`export PATH=$PATH:/usr/local/go/bin`，`export GOPROXY=https://goproxy.cn`
- 验证安装结果：`go version`

```shell
go version go1.22.4 linux/amd64
```

