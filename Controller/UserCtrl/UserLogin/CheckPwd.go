package UserLogin

func CheckPassword(sendPwd string, rightPwd string) bool {
	if sendPwd == rightPwd {
		return true
	} else {
		return false
	}
}
