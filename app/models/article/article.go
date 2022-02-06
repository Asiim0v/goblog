package article

import (
	"goblog/pkg/route"
	"strconv"
)

// Article 文章类型
type Article struct {
	ID    uint64
	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}

func (article Article) GetStringID() string {
	return strconv.FormatUint(article.ID, 10)
}