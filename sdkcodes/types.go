package sdkcodes

// Code 状态码
type Code interface {
	// Code 状态码
	Code() int64
	// Msg 消息
	Msg() string
	// I18n 国际化key
	I18n() string
	// I18nMeta 国际化扩展值
	I18nMeta() map[string]string
	/// WithMsgf 替换消息
	WithMsgf(format string, args ...any) Code
	/// WithMsgg 替换消息
	WithMsg(msg string) Code
	/// WithI18nMeta 替换国际化扩展值
	WithI18nMeta(i18nNeta map[string]string) Code
	error
}
