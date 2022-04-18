package vo

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

// ==================== COMMON STRUCTURES ====================

type ErrNo int

const (
	OK             ErrNo = 0
	ParamInvalid   ErrNo = 1 // 参数不合法
	UserHasExisted ErrNo = 2 // 该 Username 已存在
	UserHasDeleted ErrNo = 3 // 用户已删除
	UserNotExisted ErrNo = 4 // 用户不存在
	WrongPassword  ErrNo = 5 // 密码错误
	LoginRequired  ErrNo = 6 // 用户未登录

	UnknownError ErrNo = 255 // 未知错误
)

type ResponseMeta struct {
	Code ErrNo
}

// ==================== USER MANAGEMENT ====================

type RegisterRequest struct {
	Username string // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
}

type RegisterResponse struct {
	Code ErrNo
}

type LoginRequest struct {
	Username string
	Password string
}

// 登录成功后需要生成 JWT

type LoginResponse struct {
	Code ErrNo
	Data struct {
		UserID string
	}
}

type LogoutRequest struct{}

// 登出成功需要删除 JWT

type LogoutResponse struct {
	Code ErrNo
}

// ==================== PROBLEM LIST ====================

type ProblemListRequest struct {
	Page      int
	Page_Size int
}
type ProblemListResponse struct {
	Code              ErrNo
	ProblemID         int
	Status            string
	Difficulty        string
	Acceptance        string
	Global_Acceptance string
}
type ProblemDetailRequest struct {
	ProblemID int
}
type ProblemDetailResponse struct {
	Code              ErrNo
	ProblemID         int
	Description       string
	Status            string
	Difficulty        string
	Acceptance        string
	Global_Acceptance string
}

// ==================== SUBMIT ====================
type SubmissionRequest struct {
	ProblemID int
	Language  string
	Code      string
}

// ==================== RESULT ====================

type ResultRequest struct {
	ProblemID int
}
type ResultResponse struct {
	ProblemID int
	Status    string
	Language  string
	RunTime   int
}
