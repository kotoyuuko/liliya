package main

// Template is the template of source code file
type Template struct {
	Path    string
	Content string
}

var tplConfigAppIni = Template{
	Path: "src/config/app.ini",
	Content: `[APP]
TEST = "hello world"
`,
}
