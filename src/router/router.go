package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/stardemo/go-similar-keywords/src/similar"
)

func init() {
	s := g.Server()
	s.Group("/", func(baseRouterGroup *ghttp.RouterGroup) {
		baseRouterGroup.ALL("/hello", func(r *ghttp.Request) {
			r.Response.WriteExit("world")
		})
		baseRouterGroup.Group("/similar", func(similarApiGroup *ghttp.RouterGroup) {
			similarApiGroup.GET("/keywords", similar.GetSimilarKeywords)
		})
	})
}
