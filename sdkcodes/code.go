package sdkcodes

import (
	"fmt"
)

type _Code struct {
	code int64
	msg  string
	i18n string
	meta map[string]string
}

func (c *_Code) Code() int64 {
	return c.code
}
func (c *_Code) Msg() string {
	return c.msg
}
func (c *_Code) I18n() string {
	return c.i18n
}

func (c *_Code) Meta() map[string]string {
	return c.meta
}

func (c *_Code) WithMsgf(format string, args ...any) Code {
	return New(c.code, fmt.Sprintf(format, args...), c.i18n, c.meta)
}
func (c *_Code) WithMsg(s string) Code {
	return New(c.code, s, c.i18n, c.meta)
}

func (c *_Code) WithMeta(meta map[string]string) Code {
	return New(c.code, c.msg, c.i18n, meta)
}

func (c *_Code) Error() string {
	return fmt.Sprintf("状态码(%d), 消息(%s), 国际化(%s, %+v)", c.code, c.msg, c.i18n, c.meta)
}

// New 创建状态码
func New(code int64, msg string, i18n string, meta map[string]string) Code {
	return &_Code{code, msg, i18n, meta}
}

var (
	OK              = New(0, "OK", "OK", nil)
	Internal        = New(1, "服务异常", "Internal", nil)
	Service         = New(2, "服务异常", "Service", nil)
	TokenInvalid    = New(3, "令牌无效", "TokenValid", nil)
	AccessLimited   = New(4, "访问受限", "AccessLimited", nil)
	RequestNotFound = New(5, "请求资源不存在", "RequestNotFound", nil)
	ParamInvalid    = New(6, "请求参数错误", "ParamInvalid", nil)
	TooManyRequests = New(7, "请求过于频繁", "TooManyRequests", nil)
)
