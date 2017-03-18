---
title: LiteIDE 的配置文件管理
date: '2017-03-18'
description:
categories:
- blog
- liteide
tags:
- blog
- liteide
- setting
---

<!-- ## LiteIDE 的配置文件 -->

LiteIDE 的配置文件默认存储在当前用户目录中，我们也可以设置存储在程序本地目录中。通过 **选项**（ **偏好设置** ）=> **LiteApp** => **存储** _存储设置到本地ini文件_ 选项来切换默认的存储位置。

### 存储至当前用户配置
LiteIDE 的配置文件默认存储在当前用户的特定目录中，如在 macOS 下存储在当前用户目录 `/.config/liteide/liteide.ini` 文件中,
这可以保证在 LiteIDE 更新后当前用户的配置保持不变，包括用户的 GOPATH 配置以及历史目录、历史文件等信息。

### 存储至程序本地配置
我们也可以将 LiteIDE 的配置文件存储在 LiteIDE 的本地目录中，即 LiteIDE的 `share/liteapp/config/liteide.ini` 文件中，
这可以保证配置信息跟随 LiteIDE 程序启用，我们可以在自己的系统中复制多份 LiteIDE，每个都保持有自己的环境信息、GOPATH 配置、历史目录等信息，
方便多个项目独立使用。 

### 命令行操作
在 LiteIDE x31 中加入了对配置文件的命令行操作行为，包括以下三种：

* --local-setting   强制使用本地配置
* --user-setting    强制使用用户配置
* --reset-setting   重置当前配置 ( 清除配置文件 )

如果在使用 LiteIDE 过程中发现出现莫名其妙的问题或者是刚启动就报错，有可能是配置文件出错，对于这种情况，除了手工删除配置文件的方法，
我们还可以通过命令行  `>liteide --reset-setting` 来重置 LiteIDE 的当前配置文件。