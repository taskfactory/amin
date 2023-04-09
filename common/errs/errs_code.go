package errs

const (
	CodeOK      = 0
	CodeUnknown = -999 // 未知错误

	CodeDBRead  = 20000 // 读DB失败
	CodeDBWrite = 20001 // 写DB失败
)
