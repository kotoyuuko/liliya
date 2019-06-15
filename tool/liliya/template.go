package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Template is the template of source code file
type Template struct {
	Path    string
	Content string
}

func replaceTargetPath() string {
	goPath := os.Getenv("GOPATH")
	goSrcPath := path.Join(goPath, "src")
	targetPath := strings.Replace(project.Path, goSrcPath, "", 1)
	targetPath = strings.Trim(targetPath, "/")
	targetPath = path.Join(targetPath, "src")

	return targetPath
}

func createFileFromTemplate(tpl Template) error {
	filePath := path.Join(project.Path, tpl.Path)

	content := strings.Replace(tpl.Content, "github.com/kotoyuuko/liliya/src", replaceTargetPath(), -1)

	return ioutil.WriteFile(filePath, []byte(content), 0644)
}

var tplGitignore = Template{
	Path: ".gitignore",
	Content: `.DS_Store
Thumbs.db
.vscode/
*.log
*.cache
*.diff
*.exe
*.exe~
*.patch
*.swp
*.tmp
src/config/app.ini
src/log/app.log
build/
`,
}

var tplConfigAppIni = Template{
	Path: "src/config/app.ini",
	Content: `[app]
runMode = debug
encryptKey = 765a85379b6e3ed0343b68a8999ba486
timezone = Asia/Shanghai

[server]
httpAddr = 127.0.0.1:7000
readTimeout = 10
writeTimeout = 10

[database]
type = mysql
host = 127.0.0.1:3306
user = liliya
password = secret
name = liliya
`,
}

var tplDatabaseMigration = Template{
	Path: "src/database/migration.go",
	Content: `package database

import (
	"github.com/kotoyuuko/liliya/src/model"
	"github.com/kotoyuuko/liliya/src/util/dao"
)

// Migrate execute auto migration
func Migrate() {
	dao.DB.AutoMigrate(&model.User{})
}
`,
}

var tplDatabaseSeeder = Template{
	Path: "src/database/seeder.go",
	Content: `package database

import (
	"github.com/kotoyuuko/liliya/src/model"
	"github.com/kotoyuuko/liliya/src/util/dao"
)

// Seed seeds data to database
func Seed() {
	userSeeder()
}

func userSeeder() {
	db := dao.DB

	count := 0
	db.Model(&model.User{}).Count(&count)

	if count == 0 {
		user := model.User{
			Name:     "root",
			Password: "",
			Role:     "admin",
		}
		db.Save(&user)
	}
}
`,
}

var tplModelUser = Template{
	Path: "src/model/user.go",
	Content: `package model

import "github.com/kotoyuuko/liliya/pkg/model"

// User contains user information
type User struct {
	model.CommonFields
	Name     string ` + "`" + `json:"name"` + "`" + `
	Password string ` + "`" + `json:"-"` + "`" + `
	Role     string ` + "`" + `json:"role" gorm:"type:enum('user', 'admin')"` + "`" + `
}
`,
}

var tplRouterRoutes = Template{
	Path: "src/router/routes.go",
	Content: `package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kotoyuuko/liliya/src/service"
)

// Router return routes of application
func Router(r *gin.Engine) *gin.Engine {
	r.GET("/", service.Root)

	return r
}
`,
}

var tplServiceRoot = Template{
	Path: "src/service/root.go",
	Content: `package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Root is the processor for root page
func Root(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
`,
}

var tplUtilConfig = Template{
	Path: "src/util/config/config.go",
	Content: `package config

import "github.com/kotoyuuko/liliya/pkg/config"

var cfg *config.File

func init() {
	var err error
	cfg, err = config.Load("./config/app.ini")
	if err != nil {
		panic(err)
	}
}

// App returns value in app section of config file
func App(key string) *config.File {
	return cfg.Section("app").Key(key)
}

// Server returns value in server section of config file
func Server(key string) *config.File {
	return cfg.Section("server").Key(key)
}

// Database returns value in database section of config file
func Database(key string) *config.File {
	return cfg.Section("database").Key(key)
}
`,
}

var tplUtilDAO = Template{
	Path: "src/util/dao/dao.go",
	Content: `package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/kotoyuuko/liliya/pkg/database"
	"github.com/kotoyuuko/liliya/pkg/logger"
	"github.com/kotoyuuko/liliya/src/util/config"
)

// DB is the global database instance
var DB *gorm.DB

func init() {
	var err error

	dialect := config.Database("type").Default("mysql").String()
	host := config.Database("host").String()
	user := config.Database("user").String()
	password := config.Database("password").String()
	name := config.Database("name").String()
	timezone := config.App("timezone").Default("UTC").String()

	args := database.ArgsString(host, user, password, name, timezone)

	DB, err = database.Connect(dialect, args)
	if err != nil {
		logger.Warn("Cannot connect to database")
	}
}
`,
}

var tplMain = Template{
	Path: "src/main.go",
	Content: `package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kotoyuuko/liliya/pkg/logger"
	"github.com/kotoyuuko/liliya/pkg/server"
	"github.com/kotoyuuko/liliya/src/database"
	"github.com/kotoyuuko/liliya/src/router"
	"github.com/kotoyuuko/liliya/src/util/config"
)

func main() {
	runMode := config.App("runMode").Default("release").String()

	if err := logger.Init("./log/app.log", runMode); err != nil {
		panic(err)
	}

	database.Migrate()
	database.Seed()

	engine := server.Engine(router.Router, logger.Logger(), runMode)

	httpAddr := config.Server("httpAddr").Default("127.0.0.1:7000").String()
	readTimeout := config.Server("readTimeout").Default(10).Int()
	writeTimeout := config.Server("writeTimeout").Default(10).Int()
	httpServer := server.Server{
		Engine:       engine,
		Addr:         httpAddr,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	httpServer.Serve()
}
`,
}

var tplService = `package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
	})
}
`
var tplModel = `package model

import "github.com/kotoyuuko/liliya/pkg/model"

type {model} struct {
	model.CommonFields
}
`
