package similar

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/stardemo/go-similar-keywords/src/utils"
)

func GetSimilarKeywords(r *ghttp.Request) {
	req := new(ReqSimilarKeywords)
	utils.ValidationInput(r, req)
	res, err := DoGetSimilarKeywords(req)
	if err != nil {
		utils.ResponseJson(r, 500, "Failed", err.Error())
	}
	utils.ResponseJson(r, 200, "OK", res)
}
