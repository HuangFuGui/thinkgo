// 控制器模块，通过各Action调用相应业务逻辑模块，生成Json、流文件等格式数据完成请求
package controller

import (
	"encoding/json"
	"net/http"
	"log"
	"thinkgo/response"
	"thinkgo/common"
)

type BaseController struct {

}

func (this *BaseController) ReturnJsonObj(resp *response.Response, w http.ResponseWriter, r *http.Request) {
	reply, err := json.Marshal(resp)
	if err != nil {
		ctl, action := common.Com.CtlAction(r.URL.Path)
		log.Fatal(ctl + "/" + action + " reply := json.Marshal() error")
	}
	w.Header().Set("Content-Type","application/json")
	w.Write(reply)
}