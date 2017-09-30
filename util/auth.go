package util

func GetToken(authorization string) string {
	var prefix string = "Basic"
	index := len(prefix) + 1
	token := authorization[index:]
	return token
}

func CheckEqual(reqBase64Str, localBase64Str string) bool {
	if reqBase64Str == localBase64Str {
		return true
	} else {
		return false
	}
}
