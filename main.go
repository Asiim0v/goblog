package main

import (
	"fmt"
	"net/http"
)

func handlerfunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // 设置响应标头
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog! </h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
			"<a href=\"mailto:Asiim0v@example.com\">Asiim0v@example.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound) // 添加 404 状态码
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")
	}
}

func main() {
	http.HandleFunc("/", handlerfunc) // http.HandleFunc 里传参的 / 意味着 任意路径
	http.ListenAndServe(":3000", nil)
}
