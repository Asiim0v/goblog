package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"io"
	"path/filepath"
	"strings"
	"text/template"
)

// D 是 map[string]interface{} 的简写
type D map[string]interface{}

// Render 渲染通用视图
func Render(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data interface{}, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图, 把文章控制器方法内的渲染代码提取到此处, tplFiles 支持加载多个模板文件
func RenderTemplate(w io.Writer, name string, data interface{}, tplFiles ...string) {
	// 1. 设置模板相对路径
	viewDir := "resources/views/"

	// 2. 遍历传参文件列表 Slice，设置正确的路径, 支持语法糖, 例如将 articles.show 更正为 articles/show
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 3. layoutFiles 是存放所有布局模板文件的 Slice
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	// 4. 合并所有文件
	allFiles := append(layoutFiles, tplFiles...)

	// 5. 解析所有模板文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
		"RouteName2URL": route.Name2URL,
	}).ParseFiles(allFiles...)
	logger.LogError(err)

	// 6. 渲染模板
	err = tmpl.ExecuteTemplate(w, name, data)
	logger.LogError(err)
}
