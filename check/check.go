package check

import "regexp"

const (
	REGEX_MOBILE    = "^1[3-9]{1}[0-9]{9}$"
	REGEX_ID_CARD   = "^[1-9][0-9]{16}[0-9xX]$"
	REGEX_USER_NAME = "^[a-zA-Z]+[a-zA-Z0-9_-]{4,17}$" // 字母开始，由英文字母数字及下划线组成，长度5-18位字符串
	REGEX_PASSWORD  = "^[a-zA-Z0-9_-|@.]{5,18}$"
)

// CheckMobile check mobile number
func CheckMobile(mobile string) bool {
	return Check(REGEX_MOBILE, mobile)
}

// CheckIdCard check id card
func CheckIdCard(idCard string) bool {
	return Check(REGEX_ID_CARD, idCard)
}

// CheckUserName check user name
func CheckUserName(username string) bool {
	return Check(REGEX_USER_NAME, username)
}

// CheckPassword check password
func CheckPassword(password string) bool {
	return Check(REGEX_PASSWORD, password)
}

func Check(pattern, s string) bool {
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}
