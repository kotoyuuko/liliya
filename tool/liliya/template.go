package main

// Template is the template of source code file
type Template struct {
	Path    string
	Content string
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
