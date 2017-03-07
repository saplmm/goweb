package route

import (
	"gopkg.in/macaron.v1"
	"goweb/handler/adminHandler"
)

func AdminRout(m *macaron.Macaron) {
	m.Get("/admin/articles", adminHandler.ArticleManege)
	m.Get("/admin/articles-add", adminHandler.ArticleManegeAdd)
	m.Post("/admin/article-manage-add-post", adminHandler.ArticleManegeAddPost)

}