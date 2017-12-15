---
title: 解决升级 xcode 后 Qt5.6.3 不能编译的问题
date: '2017-12-15'
description:
categories:
- blog
- dev
- liteide
tags:
- blog
- liteide
---

<!-- ## 解决升级 xcode 后 Qt5.6.3 不能编译的问题 -->

升级 xcode 后，Qt5.6.3 不能编译，提示类似下列错误
```
clang: warning: no such sysroot directory: '/Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.12.sdk' [-Wmissing-sysroot]

In file included from /Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/QObject:1:
In file included from /Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/qobject.h:40:
In file included from /Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/qobjectdefs.h:41:
In file included from /Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/qnamespace.h:37:
In file included from /Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/qglobal.h:75:
/Users/user/Qt5.6.3/5.6.3/clang_64/lib/QtCore.framework/Headers/qsystemdetection.h:197:12: fatal error: 'TargetConditionals.h' file not found
#  include <TargetConditionals.h>
           ^~~~~~~~~~~~~~~~~~~~~~
1 error generated.
```

解决方法

```
$cd /Applications/Xcode.app/Contents/Developer/Platforms/MacOSX.platform/Developer/SDKs
$sudo ln -s MacOSX.sdk MacOSX10.12.sdk
```

