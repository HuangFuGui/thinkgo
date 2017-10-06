//公共模块，一些公用函数
package common

import "strings"

var Com *Common = new(Common)

type Common struct {

}

func (this *Common) CtlAction(url string) (ctl, action string) {//形如/test/hello
	path := strings.ToLower(url)
	ctlaction := strings.Split(path, "/")
	ctl = string(ctlaction[1][0] - ('a' - 'A')) + ctlaction[1][1:] + "Controller"
	action = string(ctlaction[2][0] - ('a' - 'A')) + ctlaction[2][1:] + "Action"
	return
}
