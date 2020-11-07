package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "系統內部錯誤")
	InvalidParams             = NewError(10000001, "參數錯誤")
	NotFound                  = NewError(10000002, "找無對應呼叫或服務")
	UnauthorizedAuthNotExist  = NewError(10000003, "認證失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "認證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "認證失敗，Token逾時")
	UnauthorizedTokenGenerate = NewError(10000006, "認證失敗，Token生成失敗")
	TooManyRequests           = NewError(10000007, "請求過多錯誤")
)
