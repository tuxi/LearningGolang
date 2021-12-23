# Golang 环境配置
### 安装
官网下载地址
```
http://golang.org
```
国内无妨访问的，可以去这个地址下载
```
https://studygolang.com/dl
```

下载完成后，双击pkg安装包安装即可，golang 默认安装在`/usr/local/go/`目录下

### 环境配置
安装完成后，在终端输入如下指令：
```
go version
```
如果显示`-bash: go: command not found`，说明确实go的环境，需要在`~/.bash_profile`文件中添加go的环境
```
export PATH=$PATH:/usr/local/go/bin
```
保存后，执行以下，让环境生效
```
source ~/.bash_profile
```

## 配置环境变量
> GOPATH和Go Modules相关的环境配置

- GOPATH配置
> Golang包含两个重要的环境变量：GOROOT和GOPATH，GOROOT存储了Go官方的源码和可执行文件，GOPATH存储了第三方的源码和可执行文件（自己的项目代码建议放在该目录下）。GOROOT在安装时已自动配置好，我们只需要配置GOPATH即可。

首先，创建GOPATH文件夹，打开终端：
```
mkdir -p ~/gopath/{bin,pkg,src}
```

编辑配置文件
```
vim ~/.bash_profile
```

新增以下代码
```
export GOROOT=/usr/local/go
export GOARCH=amd64
export GOOS=darwin
export GOPATH=$HOME/gopath/go
export PATH=$PATH:$GOPATH
export PATH=$PATH:$GOPATH/bin
```

保存之后，执行`source ~/.bash_profile`，运行`go env`指令即可验证GOPATH是否设置成功。我们会将GOPATH/bin文件夹加入系统环境变量，这样才能保证第三方库的可执行文件可以正常运行。

###### 注意：Mac下的 GOPATH 变量的最后一层目录名必须为 go，且该目录下必须有 src,bin,pkg 目录，否则`GOPATH`环境变量不会生效，我一开始设置的`export GOPATH=$HOME/gopath`，最后没有加go目录，导致vscode下的环境是一个新的`$HOME/go`下面，这是坑

- Go Modules配置
> 从1.11版本开始，Golang引入了新的依赖管理机制Go Modules解决长期以来Go语言依赖包没有版本控制的缺陷，Go Modules依赖的环境变量为GOPROXY和GOSUMDB，GOPROXY用于检索依赖包的信息，GOSUMDB用于校验，默认的配置为：

```
GOPROXY="https://proxy.golang.org,direct"
GOSUMDB="sum.golang.org"
```

由于国内屏蔽google，故导致这两个域名都无法访问。对于GOPROXY，七牛云做了一个镜像，方便国内开发者使用，项目地址：
[https://github.com/goproxy/goproxy.cn](https://github.com/goproxy/goproxy.cn/)

对于GOSUMDB，google官方提供了国内可访问的域名：http://sum.golang.google.cn（[参见](https://github.com/golang/go/issues/31755)）。因此，需要重新配置，同样修改配置文件
```
vim ~/.bash_profile
```
添加以下代码
```
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=sum.golang.google.cn
```

或者直接通过go指令修改：
```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```

查看go环境配置
```
go env
```

### Hello World
新建文件`main.go`
```
vim main.go
```

添加以下代码
```
// 包名,随意
package main

// 引入包
import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

保存后，运行这段代码
```
go run main.go
```

### Go语言安装包自带工具

- 在$GOROOT/bin（也就是/usr/local/go/bin） 中有三个工具
    - go 编译、运行、构建等都可以使用这个命令
    - godoc 查看包或函数的源码
    - gofmt 格式化文件

- 常用参数
    - go version 查看Go语言版本
    - go env 查看Go语言详细环境
    - go list 查看Go语言文件目录
    - go build 把源码文件构建成系统可执行的文件
    - go clean 清空生成的可执行文件
    - go vet 静态解析文件，检查是否有语法错误等
    - go get 从远程下载第三方Go语言库
    - go bug 提交bug
    - go test 测试
    - go run 编译并运行文件

### VSCode 配置Golang