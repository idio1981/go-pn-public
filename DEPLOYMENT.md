# 部署指南

## 完成GitHub推送

由于需要GitHub认证，请手动执行以下命令完成推送：

```bash
# 推送代码到GitHub
git push origin main

# 推送标签到GitHub
git push origin v1.0.0
```

## 验证模块可用性

推送完成后，其他项目可以通过以下方式引用此模块：

### 1. 直接引用最新版本
```bash
go get github.com/idio1981/go-pn-public
```

### 2. 引用特定版本
```bash
go get github.com/idio1981/go-pn-public@v1.0.0
```

### 3. 在代码中使用
```go
import (
    "github.com/idio1981/go-pn-public/api"
    "github.com/idio1981/go-pn-public/logger"
)
```

## 测试模块引用

创建一个测试项目来验证模块是否正常工作：

```bash
mkdir test-import
cd test-import
go mod init test-import
go get github.com/idio1981/go-pn-public@v1.0.0
```

然后创建main.go文件测试引用是否成功。

## 后续版本管理

当需要发布新版本时：

1. 更新代码
2. 创建新的Git标签：`git tag v1.1.0`
3. 推送代码和标签：`git push origin main --tags`
4. 更新README中的版本信息
