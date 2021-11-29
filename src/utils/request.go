// Package utils
package utils

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gvalid"
	"net/http"
)

// ValidationInput 输入校验表单
// 输入*ghttp.Request,校验结构体指针
// 出现异常返回http.status 400 /500 附带错误信息
func ValidationInput(r *ghttp.Request, pointer interface{}) {
	if err := r.Parse(pointer); err != nil {
		// Validation error.
		if v, ok := err.(gvalid.Error); ok {
			ResponseJson(r, http.StatusBadRequest, v.FirstString())
		}
		// Other error.
		ResponseJsonWithStatus(r, http.StatusInternalServerError, err.Error())
	}
}
