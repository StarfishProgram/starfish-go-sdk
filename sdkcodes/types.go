package sdkcodes

// Code 状态码
type Code interface {
	// Code 状态码
	Code() int64
	// Msg 消息
	Msg() string
	// I18n 国际化key
	I18n() string
	// Meta 国际化扩展值
	Meta() map[string]string
	/// WithMsgf 替换消息
	WithMsgf(format string, args ...any) Code
	/// WithMsgg 替换消息
	WithMsg(msg string) Code
	/// WithMeta 替换国际化扩展值
	WithMeta(meta map[string]string) Code
	error
}
