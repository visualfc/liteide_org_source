---
title: 支持 Quick Open Command
date: '2017-12-04'
description:
categories:
- blog
- dev
- liteide
tags:
- blog
- liteide
---

<!-- ## 实现 Quick Open Command -->

Sublime 和 VSCode 都有命令面板功能，方便快速调用命令。
LiteIDE 在 2016年7月发布的 X30 版本开始实现了 QuickOpen 插件，
目前支持 文件/编辑器/符号/行号 的快速跳转，但未实现菜单命令的快速跳转功能。  

在最新的开发版本中，LiteIDE 加入了菜单命令的快速跳转功能 - Quick Open Command，
快捷键为 Ctrl+Shift+P / Command+Shift+P (macos)。

如下图显示，显示当前文件可用的菜单命令。
<img src="{{urls.media}}/2017-12-04-quickopen-command/quickopencommand.png" alt="" width="600">

菜单与命令之间没有使用:分隔，而是用.来分隔，这样方便输入筛选。
<img src="{{urls.media}}/2017-12-04-quickopen-command/quickopencommand-filter.png" alt="" width="600">

回车将直接执行选中的命令。

### 实现 Quick Open Command 功能

LiteIDE 的 QuickOpen 接口 API 位于 [liteidex/src/api/quickopenapi/quickopenapi.h](https://github.com/visualfc/liteide/blob/master/liteidex/src/api/quickopenapi/quickopenapi.h) 文件中。

接口如下
```
class IQuickOpen : public QObject
{
    Q_OBJECT
public:
    IQuickOpen(QObject *parent = 0) : QObject(parent) {}
    virtual QString id() const = 0;
    virtual QString info() const = 0;
    virtual void activate() = 0;
    virtual QAbstractItemModel *model() const = 0;
    virtual void updateModel() = 0;
    virtual QModelIndex filterChanged(const QString &text) = 0;
    virtual void indexChanged(const QModelIndex &index) = 0;
    virtual bool selected(const QString &text, const QModelIndex &index) = 0;
    virtual void cancel() = 0;
};
```
其中 updateModel 函数负责更新数据，实现流程一下:

#### 获取当前所有菜单 ID 列表

```
LiteApi::IActionManager *manager = m_liteApp->actionManager();
QStringList menuIdList = manager->menuList();
("menu/build", "menu/debug", "menu/edit", "menu/file", "menu/find", "menu/help", "menu/recent", "menu/view")
```

#### 获取菜单的 Aciton 列表

根据 id 取得 QMenu

```
QMenu *menu = manager->loadMenu(idMenu);
```
注意，因为 Edit / Build 菜单都是动态加载的，直接使用 menu->action()是无法获得完整列表的。
可以通过以下代码获取。

```
QAction *menuAct = menu->menuAction();
QMenu *realMenu = menuAct->menu();
QList<QAction*> actionList = realMenu->actions();
```
#### 获取 Action 的 id, text, shortcut

在写 QuickOpenCommand 之前，大部分 QAction 是无法获取 id 的，但因为 LiteIDE 中的 Action 通过

IActionManager 注册以支持 自定义快捷键功能，所以在通过 IActionContext::regAction 函数注册 Aciton 时的 id/shortcut 时加入 act->setData(id) 将 Aciton 注册用 id 保存下来，这样就可以通过 act->data() 可以获取到 id, act->text() 获取翻译后的文本，act->shortcut() 获取 快捷键。

将获取的 id,text,shortcut 加入列表中即可。

#### 注册

通过 IQuickOpenManager::addFilter(const QString &sym, IQuickOpen *filter) 函数可以注册 QuickOpen，QuickOpenCommand 的 sym 选用的是 > ,在 view 菜单中加入相应 Action，快捷键是 Ctrl+Shift+P / Command+Shift+P 。

#### 实现代码

实现代码位于 [liteidex/src/plugins/quickopen](https://github.com/visualfc/liteide/tree/master/liteidex/src/plugins/quickopen)
目录下 quickopenation.h 和 quickopenaction.cpp 文件中。




