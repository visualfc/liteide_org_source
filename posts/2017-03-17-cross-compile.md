---
title: cross-compile
date: '2017-03-17'
description:
categories:
- blog
- liteide
tags:
- blog
- liteide
- cross-compile
---
## LiteIDE 的交叉编译设置

LiteIDE 通过使用环境配置文件来支持 go 语言的交叉编译，对于 windows 和 linux 下相对简单，对于 macOS 可能会复杂一些，
本文以 macOS 交叉编译 windows-386 可执行文件为例，简要介绍如何在 LiteIDE 中配置和实现交叉编译功能。


### 交叉编译环境选择和配置

首先选择环境，在工具栏环境选择中选择  cross-win32 切换至 windows-386 交对编译环境，我们点击工具栏上的 **编辑当前环境** 按钮，对当前使用的环境即 cross-win32.env 文件进行编辑，默认设置如下：


	GOROOT=$HOME/go
	#GOBIN=
	GOARCH=386
	GOOS=windows
	CGO_ENABLED=0
	
	PATH=$GOROOT/bin:$PATH
	
	LITEIDE_GDB=/usr/local/bin/gdb
	LITEIDE_MAKE=make
	LITEIDE_TERM=/usr/bin/open
	LITEIDE_TERMARGS=-a Terminal
	LITEIDE_EXEC=/usr/X11R6/bin/xterm
	LITEIDE_EXECOPT=-e


我们可以看到 GOARCH 和 GOOS 已经设置完毕，我们需要更改的是 GOROOT 和 PATH 变量，主要设置方式有两种，

#### 第一种方式，不设置 GOROOT 直接设置 PATH

这种方式中，我们不设置 GOROOT 变量，而是通过 PATH 设置让 LiteIDE 在内部使用 go env 自动查询 GOROOT，设置如下

```
#GOROOT=$HOME/go
#GOBIN=
GOARCH=386
GOOS=windows
CGO_ENABLED=0

PATH=$PATH:/usr/local/bin

LITEIDE_GDB=/usr/local/bin/gdb
LITEIDE_MAKE=make
LITEIDE_TERM=/usr/bin/open
LITEIDE_TERMARGS=-a Terminal
LITEIDE_EXEC=/usr/X11R6/bin/xterm
LITEIDE_EXECOPT=-e
```
#### 第二种方式，直接设置 GOPATH 值 
这种方式是直接设置 GOROOT，可以通过在终端下输入 `go env` 来查询 GOROOT 位置，如果是使用官方安装，
可能显示为
```
GOROOT="/usr/local/go"
```
如果是使用 brew install go 安装，则可能显示为
```
GOROOT="/usr/local/Cellar/go/1.8/libexec"
```
对于官方安装，我们编辑 cross-win32 环境对应为
```
GOROOT=/usr/local/go
#GOBIN=
GOARCH=386
GOOS=windows
CGO_ENABLED=0

PATH=$GOROOT/bin:$PATH

LITEIDE_GDB=/usr/local/bin/gdb
LITEIDE_MAKE=make
LITEIDE_TERM=/usr/bin/open
LITEIDE_TERMARGS=-a Terminal
LITEIDE_EXEC=/usr/X11R6/bin/xterm
LITEIDE_EXECOPT=-e
```
对于 brew 安装，我们编辑 cross-win32 环境对应为
```
GOROOT=/usr/local/Cellar/go/1.8/libexec
#GOBIN=
GOARCH=386
GOOS=windows
CGO_ENABLED=0

PATH=$GOROOT/bin:$PATH

LITEIDE_GDB=/usr/local/bin/gdb
LITEIDE_MAKE=make
LITEIDE_TERM=/usr/bin/open
LITEIDE_TERMARGS=-a Terminal
LITEIDE_EXEC=/usr/X11R6/bin/xterm
LITEIDE_EXECOPT=-e
```

#### 保存环境并通过 go env 查看
保存cross-win32.env 后，LiteIDE 自动重新加载 cross-win32 环境，打开编译输出窗口可以看到类似如下输出
```
/usr/local/go/bin/go env []
GOARCH="386"
GOBIN=""
GOEXE=".exe"
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOOS="windows"
GOPATH="/Users/vfc/goproj"
GORACE=""
GOROOT="/usr/local/go"
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
GO386=""
CC="clang"
GOGCCFLAGS="-m32 -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/2b/fhf209x571lgp5dljz593q040000gn/T/go-build420337682=/tmp/go-build -gno-record-gcc-switches"
CXX="clang++"
CGO_ENABLED="0"
PKG_CONFIG="pkg-config"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
```
注：以上环境信息也可以调用菜单 **查看** - **执行文件** （快捷键为 Command+` ）功能的输入窗口中输入 go env 来查看。

### 交叉编译时提示没有权限的处理方法

在 cross-win32环境下，我们打开项目中的 go 源码文件，使用工具栏上的 **编译** 命令或者在 **执行文件** 中输入 go build 进行编译测试。
如果是使用 brew 安装，这时候就可以编译成功，如果使用 go 安装，有可能会出现以下提示 
```
go install runtime/internal/sys: mkdir /usr/local/go/pkg/windows_386: permission denied
```
出现这个错误则说明 LiteIDE 没有 /usr/local/go 路径的写权限，遇到这种情况，我们需要在终端下先交叉编译好 go 环境。在终端下输入
```
$ GOOS=windows GOARCH=386 go install std
```
这里的 `go install std` 表示编译安装标准库，安装结束后，我们在 LiteIDE 中就可以正常使用 windows-386 交叉编译环境了。
对于其他交叉编译环境，需要时可以做同样的处理。