package controller

import (
	"net/http"
	"thinkgo/router"
	"thinkgo/service"
	"thinkgo/response"
)

type TestController struct {
	BaseController
}

func init(){
	router.RegisterController(&TestController{})
}

func (this *TestController) HelloAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.HelloService(resp, r)
	this.ReturnJsonObj(resp, w, r)
}

func (this *TestController) LoginAction(w http.ResponseWriter, r *http.Request){
	svc, resp := new(service.TestService), new(response.Response)
	svc.LoginService(resp, r)
	this.ReturnJsonObj(resp, w, r)
}