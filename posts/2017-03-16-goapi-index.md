---
title: goapi-index
date: '2017-03-16'
description:
categories:
- blog
- liteide
tags:
- goapi
---
最近更新了 Golang Api Index 功能，之前是通过 gotools 对 GOPATH 内的所有源码进行索引，并在用户目录中建立一个索引文件，
文件格式与 go/api/go1.txt 类似，这个功能的问题是使用 gotools 建立索引的速度很慢，查询功能与 Golang Doc Search 的功能也有所重复，而且我还发现 gotools 工具对于 go1.8 的源码建立索引出错。

在最新版本中，我对 Golang Api Index 的功能重新定位，更改为只查询 go 标准库 api，即查询 go/api 目录下的文件，
包含 `go1.txt go1.1.txt ... except.txt next.txt` ，
同时在查询结果列表中加入了 api 所对应的 go 版本（文件），如下图所示。

<img src="{{urls.media}}/2017-03-16-goapi-index/goapi-index.png" alt="" width="600">

这样看起来工整多了，并且与 Golang Doc Search 有所区别。