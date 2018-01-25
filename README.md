# importgo

一步一步学习 go 语言基础。

* 视频教程：https://space.bilibili.com/276273794/#/
* Gitbook 电子书：https://songjiayang.gitbooks.io/go-basic-courses/content/
* Go 实战交流 QQ 群：694650181 （申请加入请备注“禾木课堂”）

### 写在前面

* 教学 Go 版本 1.9.x
* 教学使用 `GOPATH` 为 `~/importgo`

### 安装

1. 下载安装对应版本 https://golang.org/dl/
2. 查看 go 安装目录 `/usr/local/go` (Windows 下默认安装到 `c:\Go`)
3. 运行 `go version` 命令检查是否安装正确

ps： 推荐使用二进制默认安装方式

### 项目环境变量

本项目命名为 `importgo`, 故添加一个 `IMPORTGOROOT` 的环境变量进行所有的代码开发和演示：

```
export IMPORTGOROOT=$HOME/importgo
export GOPATH=$IMPORTGOROOT # 覆盖 GOPATH 环境变为 importgo
export PATH=$IMPORTGOROOT/bin:$PATH #
```
