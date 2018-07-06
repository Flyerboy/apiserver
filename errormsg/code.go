package errormsg

/*
错误码有3部分组成，一共 5 位
第一位： 1.系统错误  2.用户错误
第2-3位：模块错误
第4-5位：具体错误
*/

var (
	OK = &ErrorMsg{Code: 0, Message: "ok"}

	// system error
	ServerError = &ErrorMsg{Code:10001, Message: "Server Error"}
	ParamsError = &ErrorMsg{Code: 10002, Message: "Params Error"}

	// user error
	UserNotFound = &ErrorMsg{Code:20001, Message: "user not found"}
)
