# Liliya

## 介绍

Liliya 是一套 Go 微服务框架，基于 Gin 和 GORM。包含 `liliya` 命令行工具，可以快速生成目录结构和代码。

## 快速开始

### 设置环境变量

先在 `.zshrc` 里写入：

    export GO111MODULE=on

### 获取 liliya 命令行工具

    go install github.com/kotoyuuko/liliya/tool/liliya

### 创建项目

> 必须在 `$GOPATH/src` 的子目录使用，否则会出错。

    cd $GOPATH/src/github.com/kotoyuuko
    liliya create liliya-demo

### 生成 Service

    cd liliya-demo
    liliya make service hello

### 生成 Model

    cd liliya-demo
    liliya make model test

## 使用文档

### 路由

修改 `src/router/routes.go` 即可，写法与 Gin 框架内置路由写法相同。

### 配置

编辑 `src/config/app.ini` 或 `build/config/app.ini`。

程序内使用配置：

    import "{project_path}src/util/config"

    config.App("runMode").String()

可参照 `src/util/config/config.go` 已有的函数增加其他函数方便使用。

### Service

基本格式：

    package service

    import (
        "net/http"

        "github.com/gin-gonic/gin"
    )

    func Root(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "status": http.StatusOK,
        })
    }

完全是 Gin Context 的写法，高级功能参照 Gin 官方手册。

### Model

基本格式：

    package model

    import "github.com/kotoyuuko/liliya/pkg/model"

    type User struct {
        model.CommonFields
        Name     string `json:"name"`
        Password string `json:"-"`
        Role     string `json:"role" gorm:"type:enum('user', 'admin')"`
    }

Liliya 使用了 GORM，可以通过下面方法获取到数据库实例：

    import "{project_path}src/util/dao"

    func main() {
        db := dao.DB
    }

具体的数据库操作参照 GORM 官方文档。

### Log

Liliya 使用 logrus 处理日志，程序内调用方法如下：

    import "github.com/kotoyuuko/liliya/pkg/logger"

    logger.Info("test")

## License

The Unlicense
