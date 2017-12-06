---
title: 支持 gomodifytags
date: '2017-12-02'
description:
categories:
- blog
- dev
- liteide
tags:
- blog
- liteide
---

<!-- ## 实现 gomodifytags -->

在 LiteIDE 最新的开发版本中，加入了对 [gomodifytags](https://github.com/fatih/gomodifytags) 的支持，
gomodifytags 的主要功能是对 Go语言结构体字段 (go struct field) 的 tags 进行增删和修改，比如 json, xml 以及其他自定义 tag 。

### 结构体的选择

* 当光标于结构内部时，Add Tags 和 Remove Tags 功能对应于当前光标下的整个结构体进行操作。
* 当光标对结构体中字段进行选择时，Add Tags 和 Remove Tags 功能对于于光标所选中的行范围内字段进行操作。

在 LiteIDE 中的 Add Tags Dialog / Remove Tags Dialog 中会有相应的提示。

### 为结构体字段增加和修改 Tags

以下面结构体为例
```
type Server struct {
	Name        string
	Port        int
	EnableLogs  bool
	BaseDomain  string
	Credentials struct {
		Username string
		Password string
	}
}
```
在 LiteIDE 中将光标置于 Name 到 } 结尾之间任意一处， 编辑菜单或右键调用 Add Tags 将对 Server 结构体全部字段进行操作。
如下图所示，选中 json 和 xml :
<img src="{{urls.media}}/2017-12-02-gomodifytags/goaddtags1.png" alt="" width="600">

确认将获得下面结果

```
type Server struct {
	Name        string `json:"name" xml:"name"`
	Port        int    `json:"port" xml:"port"`
	EnableLogs  bool   `json:"enable_logs" xml:"enable_logs"`
	BaseDomain  string `json:"base_domain" xml:"base_domain"`
	Credentials struct {
		Username string `json:"username" xml:"username"`
		Password string `json:"password" xml:"password"`
	} `json:"credentials" xml:"credentials"`
}
```

json 和 xml 都可以自定义选项，多个选项以 , 作为分隔。

### 删除结构体字段的 Tags

如果要删除上例中的 xml tags，同样将光标位于 Name 到 } 之间任意一处，编辑菜单或右键调用 Remove Tags 命令进行操作。
如下图所示，选中 json
<img src="{{urls.media}}/2017-12-02-gomodifytags/goremovetags1.png" alt="" width="600">

确认将删除 xml tags， 所下图所示。

```
type Server struct {
	Name        string `json:"name"`
	Port        int    `json:"port"`
	EnableLogs  bool   `json:"enable_logs"`
	BaseDomain  string `json:"base_domain"`
	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}
```

删除也可以删除 tags 对应的选项，根据需要操作即可。