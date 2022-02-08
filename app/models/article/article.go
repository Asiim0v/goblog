package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"strconv"
)

// Article 文章类型
type Article struct {
	models.BaseModel

	Title string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body  string `gorm:"type:longtext;not null;" valid:"body"`
}

// Link 方法用来生成文章链接
func (article Article) Link() string {
	return route.Name2URL("articles.show", "id", article.GetStringID())
}

func (article Article) GetStringID() string {
	return strconv.FormatUint(article.ID, 10)
}
