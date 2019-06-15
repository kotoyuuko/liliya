package main

import (
	"io/ioutil"
	"path"
)

// Template is the template of source code file
type Template struct {
	Path    string
	Content string
}

func createFileFromTemplate(tpl Template) error {
	filePath := path.Join(project.Path, tpl.Path)

	return ioutil.WriteFile(filePath, []byte(tpl.Content), 0644)
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
