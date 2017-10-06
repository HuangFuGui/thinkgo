//路由模块，负责路由http分发请求至相应controller/action
package router

import (
	"net/http"
	"reflect"
	"thinkgo/common"
)

var Routermap map[string]interface{} = make(map[string]interface{})

func RegisterController(i interface{}) {
	value := reflect.ValueOf(i)
	ctl := value.Type().Elem().Name()
	Routermap[ctl] = value
}

func HttpRouter(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	url := r.URL.Path

	//方式一：硬编码（但直接这么写容易造成代码结构循环引用）
	//可脚本自动化生成，结合IOC容器的概念，程序启动时即装载到map等容器，不走反射，提高路由效率，TODO
	//if url == "/test/hello" {
	//	ctl := new(controller.TestController)
	//	ctl.HelloAction(w, r)
	//}
	//if url == "/test/login" {
	//	ctl := new(controller.TestController)
	//	ctl.LoginAction(w, r)
	//}

	//方式二：反射
	ctl, action := common.Com.CtlAction(url)
	args := make([]reflect.Value, 2)
	args[0] = reflect.ValueOf(w)
	args[1] = reflect.ValueOf(r)
	Routermap[ctl].(reflect.Value).MethodByName(action).Call(args)
}