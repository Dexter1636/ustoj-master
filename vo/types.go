package vo

import "ustoj-master/model"

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
	Username string `form:"username"` // required，只支持大小写，长度不小于 8 位 不超过 20 位
	Password string `form:"password"` // required，同时包括大小写、数字，长度不少于 8 位 不超过 20 位
}

type RegisterResponse struct {
	Code ErrNo `json:"code"`
}

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

// 登录成功后需要生成 JWT

type LoginResponse struct {
	Code ErrNo `json:"code"`
	Data struct {
		UserID string `json:"user_id"`
	}
	Token string `json:"token"`
}

type LogoutRequest struct{}

// 登出成功需要删除 JWT

type LogoutResponse struct {
	Code ErrNo `json:"code"`
}

// ==================== PROBLEM LIST ====================

type ProblemListRequest struct {
	Page      int `json:"page"`
	Page_Size int `json:"page_size"`
}
type ProblemListResponse struct {
	Code        ErrNo           `json:"code"`
	Problemlist []model.Problem `json:"problemlist"`
	Username    string          `json:"username"`
}
type ProblemDetailRequest struct {
	ProblemID int `json:"problem_id"`
}
type ProblemDetailResponse struct {
	Code              ErrNo  `json:"code"`
	ProblemID         int    `json:"problem_id"`
	Description       string `json:"description"`
	Status            string `json:"status"`
	Difficulty        string `json:"difficulty"`
	Acceptance        string `json:"acceptance"`
	Global_Acceptance string `json:"global_acceptance"`
	Username          string `json:"username"`
}

// ==================== SUBMIT ====================
type SubmissionRequest struct {
	ProblemID int    `json:"problem_id"`
	Language  string `json:"language"`
	Code      string `json:"code"`
}

type SubmissionResponse struct {
	Code ErrNo `json:"code"`
}

// ==================== RESULT ====================

type ResultRequest struct {
	ProblemID int    `json:"problem_id"`
	Username  string `json:"username"`
}
type ResultResponse struct {
	Code      ErrNo  `json:"code"`
	ProblemID int    `json:"problem_id"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	Language  string `json:"language"`
	RunTime   int    `json:"run_time"`
}
