package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

// Render 渲染视图, 把文章控制器方法内的渲染代码提取到此处
func Render(w io.Writer, name string, data interface{}) {
	// 1. 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 语法糖, 将 articles.show 更正为 articles/show
	name = strings.Replace(name, ".", "/", -1)

	// 3. files 是存放所有布局模板文件的 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	// 4. 在 Slice 中新增我们的目标文件
	newFiles := append(files, viewDir+name+".gohtml")

	// 5. 解析所有模板文件
	tmpl, err := template.New(name + ".gohtml").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(newFiles...)
	logger.LogError(err)

	// 6. 渲染模板
	err = tmpl.ExecuteTemplate(w, "app", data)
	logger.LogError(err)
}