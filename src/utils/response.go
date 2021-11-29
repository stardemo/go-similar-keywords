package utils

/*
   @Time : 2020/7/2 1:45 下午
   @Author : starliu
   @File : response.go
*/
import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

// ResponseJsonWithStatus 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// status:  Http Status;
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func ResponseJsonWithStatus(r *ghttp.Request, status int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteHeader(status)
	err := r.Response.WriteJsonExit(g.Map{
		"rspcode": status,
		"rspdesc": msg,
		"data":    responseData,
	})
	if err != nil {
		g.Log().Critical("HTTP数据返回错误", err)
		r.Response.WriteHeader(500)
		r.Response.Write("Internal Server Error!")
	}
}

// ResponseJson 标准返回结果数据结构封装。
// 返回固定数据结构的JSON:
// code:  错误码(0:成功, 1:失败, >1:错误码);
// msg:  请求结果信息;
// data: 请求结果,根据不同接口返回结果的数据结构不同;
func ResponseJson(r *ghttp.Request, code int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteHeader(http.StatusOK)
	err := r.Response.WriteJsonExit(g.Map{
		"rspcode": code,
		"rspdesc": msg,
		"data":    responseData,
	})
	if err != nil {
		g.Log().Critical("HTTP数据返回错误", err)
		r.Response.WriteHeader(500)
		r.Response.Write("Internal Server Error!")
	}
}

// ResponseStringWithStatus 标准返回结果数据结构封装。
// 返回状态码以及字符串:
func ResponseStringWithStatus(r *ghttp.Request, status int, data string) {
	r.Response.WriteHeader(status)
	r.Response.WriteExit(data)
}

// ResponseString 标准返回结果数据结构封装。
// 返回状态码以及字符串:
func ResponseString(r *ghttp.Request, data string) {
	r.Response.WriteHeader(http.StatusOK)
	r.Response.WriteExit(data)
}
