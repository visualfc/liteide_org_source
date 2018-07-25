---
title : LiteIDE 简介
description:
---

LiteIDE X
=========

### 简介

_LiteIDE是一个简单，开源，跨平台的Go语言IDE._

* 版本: X34
* 作者: [七叶 (visualfc)](mailto:visualfc@gmail.com)


### 功能
* 核心功能
	* 系统环境管理
	* 编辑命令配置
	* 简单开放的调试系统
	* Kate格式语法高亮和自动完成
	* 可配置的语法WordApi列表
	* 基于Mime类型的系统
	* 完全插件支持
* Golang语言支持
	* 包浏览器
	* 结构视图和大纲
	* 文档浏览
	* [Gocode](https://github.com/nsf/gocode)代码自动完成支持
	* Go API检索
	* 代码导航
	* 查找使用
	* 代码重构
	* Go playground
* 其他支持
	* Markdown
	* Json
	* Golang Present	

### 系统支持
* Windows x86 (32-bit or 64-bit) 
* Linux x86 (32-bit or 64-bit)
* MacOS X10.6

### LiteIDE 命令行
	liteide [files|folder] [--select-env id] [--local-setting] [--user-setting] [--reset-setting]
	 
	--select-env [system|win32|cross-linux64|...]     选择初始环境ID
	--local-setting   强制使用本地配置
	--user-setting    强制使用用户配置
	--reset-setting   重置当前配置 ( 清除配置文件 )

### 网址
* 源代码下载
	* <https://github.com/visualfc/liteide>
* 发行版下载
	* <http://sourceforge.net/projects/liteide/files>
* Google用户组
	* <https://groups.google.com/group/liteide-dev>
* FAQ
	* <https://github.com/visualfc/liteide/blob/master/liteidex/deploy/welcome/zh_CN/guide.md>	
* 支持LiteIDE
	* <http://visualfc.github.com/support>
