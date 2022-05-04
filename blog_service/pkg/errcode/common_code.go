package errcode

var (
	Success =NewError(0,"成功")
	ServerError = NewError(1000000,"服务内部错误")
	InvaildParams = NewError(10000001,"入参错误")
	NotFound = NewError(10000002,"找不到")
	UnauthorizedAuthNotExist =NewError(10000003,"鉴权失败，找不到对应的APPKey和AppSecret")
	UnauthorizedTokenError =NewError(10000004,"鉴权失败，token错误")
	UnauthorizedTokenTimeout =NewError(10000005,"鉴权失败，token超时")
	UnauthorizedTokenGenerate =NewError(10000006,"鉴权失败，token生成失败")
	TooManyRequests = NewError(10000007,"请求过多")
)
