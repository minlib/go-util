package check

import "regexp"

// CheckMobile check mobile number
func CheckMobile(mobile string) bool {
	return Check("^1[3-9]{1}[0-9]{9}$", mobile)
}

// CheckIdCard check id card
func CheckIdCard(idCard string) bool {
	return Check("^[1-9][0-9]{16}[0-9xX]$", idCard)
}

// CheckUserName check user name
func CheckUserName(username string) bool {
	return Check("^[a-zA-Z]+[a-zA-Z0-9_-]{4,17}$", username)
}

// CheckPassword check password
func CheckPassword(password string) bool {
	return Check("^[a-zA-Z0-9_-|@.]{5,18}$", password)
}

func Check(pattern, s string) bool {
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}
