package adminHandler

import "gopkg.in/macaron.v1"

func ArticleManege(ctx *macaron.Context) {
	ctx.HTML(200, "admin/article-manage")
}

func ArticleManegeAdd(ctx *macaron.Context) {
	ctx.HTML(200, "admin/article-manage-add")
}

func ArticleManegeAddPost(ctx *macaron.Context) {

}
