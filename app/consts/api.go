package consts

const (
	Fail    = 100
	ERROR   = 500
	SUCCESS = 200
)

var MsgFlags = map[int]string{
	Fail:    "fail",
	SUCCESS: "success",
	ERROR:   "error",
}

func GetMessage(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
