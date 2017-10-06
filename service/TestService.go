package service

import (
	"net/http"
	"thinkgo/response"
)

type TestService struct {
	BaseService
}

func (this *TestService) HelloService(resp *response.Response, r *http.Request) {
	resp.Code = 0
	resp.Msg = "OK"
	resp.Data = "hello thinkgo~"
}

func (this *TestService) LoginService(resp *response.Response , r *http.Request) {
	if len(r.Form["username"]) < 1 || len(r.Form["username"][0]) == 0 {
		resp.Code = -1
		resp.Msg = "用户名" + response.SHOULDNOTEMPTY
		resp.Data = nil
	}else if len(r.Form["password"]) < 1 || len(r.Form["password"][0]) == 0 {
		resp.Code = -1
		resp.Msg = "密码" + response.SHOULDNOTEMPTY
		resp.Data = nil
	}else {
		resp.Code = 0
		resp.Msg = "OK"
		data := "用户名为：" + r.Form["username"][0]+"，密码为：" + r.Form["password"][0]
		resp.Data = data
	}
}