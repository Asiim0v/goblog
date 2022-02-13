package controllers

import (
	"goblog/app/models/article"
	"goblog/app/models/user"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// UserController 用户控制器
type UserController struct {
	BaseController
}

// Show 用户个人页面
func (uc *UserController) Show(w http.ResponseWriter, r *http.Request) {

	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的数据
	_user, _ := user.Get(id)

	// 3. 获取结果集
	articles, pagerData, err := article.GetByUserID(_user.GetStringID(), r, 2)

	if err != nil {
		uc.ResponseForSQLError(w, err)
	} else {
		// ---  4. 加载模板 ---
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}
