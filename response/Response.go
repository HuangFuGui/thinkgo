package response

const (
	SHOULDNOTEMPTY	=	"不能为空"
	ISWRONG			=	"错误"
)

type Response struct {
	Code 	int
	Msg 	string
	Data 	interface{}
}