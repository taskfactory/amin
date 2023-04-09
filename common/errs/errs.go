package errs

import "fmt"

// Err 自定义错误
type Err struct {
	Code int32
	Msg  string
}

// New 构造函数
func New(code int32, msg string) *Err {
	return &Err{Code: code, Msg: msg}
}

// Newf 构造函数
func Newf(code int32, format string, args ...interface{}) *Err {
	return &Err{Code: code, Msg: fmt.Sprintf(format, args...)}
}

// Error 实现errors.Error接口
func (e *Err) Error() string {
	return fmt.Sprintf("code:%v,msg:%s", e.Code, e.Msg)
}

// Code 获取错误码
func Code(e error) int32 {
	if e == nil {
		return CodeOK
	}
	err, ok := e.(*Err)
	if ok {
		return err.Code
	}

	return CodeUnknown
}

// Msg 获取错误信息
func Msg(e error) string {
	if e == nil {
		return ""
	}
	err, ok := e.(*Err)
	if ok {
		return err.Msg
	}

	return e.Error()
}
