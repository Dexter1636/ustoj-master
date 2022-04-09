package vo

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

// 通用结构

type ErrNo int

const (
	OK           ErrNo = 0
	ParamInvalid ErrNo = 1 // 参数不合法

	UnknownError ErrNo = 255 // 未知错误
)

type ResponseMeta struct {
	Code ErrNo
}
